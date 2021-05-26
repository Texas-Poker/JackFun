package redis_helper

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
	"server/system"
	"strings"
	"time"
)

type (
	// AsyncPipelineType is the type of async pipeline instance
	AsyncPipelineType = string

	asyncPipelineOptions struct {
		clientTyp string
		parallel  int
		tick      int
		queued    int
	}

	asyncPipelineCmd struct {
		cmders   []redis.Cmder
		callback func([]redis.Cmder) error
	}

	asyncPipeline struct {
		pipeliner redis.Pipeliner
		chQueue   chan *asyncPipelineCmd
		queued    int
		chDie     chan struct{}
		chExit    chan struct{}
	}

	asyncPipelinePool struct {
		pipelines []*asyncPipeline
		options   *asyncPipelineOptions
	}

	// AsyncPipeliner is an interface for asyncPipeline
	AsyncPipeliner interface {
		DoCmd(cmder redis.Cmder) error
		DoCmds(cmders []redis.Cmder) error
		DoCmdWithCallback(cmder redis.Cmder,
			callback func(redis.Cmder) error) error
		DoCmdsWithCallback(cmders []redis.Cmder,
			callback func([]redis.Cmder) error) error
	}
)

var (
	asyncPipelinePools = make(map[string]*asyncPipelinePool)
)

// DoCmd is a thread-safe func.
// It pushes cmder into queue, and exec it async.
func (p *asyncPipeline) DoCmd(cmder redis.Cmder) error {
	if cmder == nil {
		return fmt.Errorf("Cmder passing in is nil")
	}

	cmders := make([]redis.Cmder, 0)
	cmders = append(cmders, cmder)
	p.chQueue <- &asyncPipelineCmd{
		cmders:   cmders,
		callback: nil,
	}
	return nil
}

// DoCmds is a thread-safe func.
// It pushes cmders into queue, and exec it async.
func (p *asyncPipeline) DoCmds(cmders []redis.Cmder) error {
	for _, cmder := range cmders {
		if cmder == nil {
			return fmt.Errorf("Cmder passing in is nil")
		}
	}

	p.chQueue <- &asyncPipelineCmd{
		cmders:   cmders,
		callback: nil,
	}
	return nil
}

// DoCmd is a thread-safe func.
// It pushes cmders into queue, and exec it async.
// Additionally, it will callback after exec.
func (p *asyncPipeline) DoCmdWithCallback(cmder redis.Cmder,
	callback func(redis.Cmder) error) error {

	if cmder == nil {
		return fmt.Errorf("Cmder passing in is nil")
	}

	cmders := make([]redis.Cmder, 0)
	cmders = append(cmders, cmder)
	p.chQueue <- &asyncPipelineCmd{
		cmders: cmders,
		callback: func(cmders []redis.Cmder) error {
			for _, cmder := range cmders {
				if callback != nil {
					if err := callback(cmder); err != nil {
						return err
					}
				}
			}
			return nil
		},
	}
	return nil
}

// DoCmdsWithCallback is a thread-safe func.
// It pushes cmders into queue, and exec it async.
// Additionally, it will callback after exec.
func (p *asyncPipeline) DoCmdsWithCallback(cmders []redis.Cmder,
	callback func([]redis.Cmder) error) error {

	for _, cmder := range cmders {
		if cmder == nil {
			return fmt.Errorf("Cmder passing in is nil")
		}
	}
	p.chQueue <- &asyncPipelineCmd{
		cmders:   cmders,
		callback: callback,
	}

	return nil
}

// Close is a thread-safe func.
// It executes all queued cmd and stop cmder queue
func (p *asyncPipeline) Close() {
	close(p.chDie)
	<-p.chExit
}

func (p *asyncPipeline) process(pCmd *asyncPipelineCmd) {
	for _, cmder := range pCmd.cmders {
		if err := p.pipeliner.Process( cmder); err != nil {
			// Just a cmd error, continuing.
			logger.Log.Errorf("Redis cmder %v error: %v", cmder.Args, err)
		} else {
			p.queued++
		}
	}
}

func (p *asyncPipeline) exec() {
	cmders, err := p.pipeliner.Exec()
	if err != nil {
		for _, cmder := range cmders {
			if cmder.Err() != nil {
				// Just a cmd error, continuing.
				logger.Log.Errorf("Redis cmder %v error: %v", cmder.Args(), cmder.Err())
			}
		}
	}
	p.queued = 0
}

func (p *asyncPipeline) callback(pCmd *asyncPipelineCmd) {
	go system.SlightPanic(func() {
		if err := pCmd.callback(pCmd.cmders); err != nil {
			// Just a cmd error, continuing.
			logger.Log.Errorf("Redis cmder callback error: %v", err)
		}
	})
}

func (p *asyncPipeline) start(options *asyncPipelineOptions) {
	go system.SlightPanic(func() {
		tick := time.NewTicker(time.Duration(options.tick) * time.Millisecond)
		defer func() {
			close(p.chQueue)
			tick.Stop()
			close(p.chExit)
		}()
		for {
			select {
			case <-p.chDie:
				system.SlightPanic(func() {
					if p.queued > 0 {
						p.exec()
					}
				})
				return
			default:
				system.SlightPanic(func() {
					select {
					case pCmd := <-p.chQueue: // It comes a cmder.
						if pCmd.callback == nil { // It comes a no-reply cmder.
							p.process(pCmd)

							// If cmder in queue is over `options.queued` limit.
							// exec queued cmders immediately.
							if p.queued >= options.queued {
								p.exec()
							}
						} else { // It comes a reply cmder
							// If cmder in queue is over `options.queued` limit,
							// exec queued cmders immediately.
							if p.queued > 0 {
								p.exec()
							}

							// Process reply cmder right now.
							p.process(pCmd)

							// Exec reply cmder right now.
							p.exec()

							// Call callback to handle cmders.
							p.callback(pCmd)
						}
					case <-tick.C:
						// After every `tick` milliseconds,
						// exec queued cmders immediately.
						if p.queued > 0 {
							p.exec()
						}
					}
				})
			}
		}
	})
}

// GetAsyncPipeliner return a async pipeliner.
func GetAsyncPipeliner(typ AsyncPipelineType, index int64) (AsyncPipeliner, error) {
	pool, err := getAsyncPipelinePool(typ)
	if err != nil {
		return nil, err
	}
	pipelineIndex := index % int64(pool.options.parallel)
	pipeline := pool.pipelines[pipelineIndex]
	return pipeline, nil
}

func getAsyncPipelinePool(typ AsyncPipelineType) (*asyncPipelinePool, error) {
	pool, ok := asyncPipelinePools[typ]
	if ok {
		return pool, nil
	}

	InitAsyncPipelinePool(typ)

	pool, ok = asyncPipelinePools[typ]
	if !ok {
		return nil, fmt.Errorf("cannot find redis async pipeline type %s", typ)
	}
	return pool, nil
}

func newAsyncPipelineOptions(typ AsyncPipelineType) *asyncPipelineOptions {



	ClientTypeKey:= configStr("pipeline", typ, "redis")

	clientTyp := pitaya.GetConfig().GetString(ClientTypeKey)
	parallel := pitaya.GetConfig().GetInt(configStr("pipeline", typ, "parallel"))
	tick := pitaya.GetConfig().GetInt(configStr("pipeline", typ, "tick"))
	queued := pitaya.GetConfig().GetInt(configStr("pipeline", typ, "queued"))

	return &asyncPipelineOptions{
		clientTyp: clientTyp,
		parallel:  int(parallel),
		tick:      int(tick),
		queued:    int(queued),
	}
}

func newAsyncPipeline(options *asyncPipelineOptions) *asyncPipeline {
	clientPool, err := getClientPool(options.clientTyp)
	if err != nil {
		panic(err)
	}
	p := &asyncPipeline{
		pipeliner: clientPool.client.Pipeline(),
		chQueue:   make(chan *asyncPipelineCmd, options.queued*2),
		queued:    0,
		chDie:     make(chan struct{}),
		chExit:    make(chan struct{}),
	}
	p.start(options)
	return p
}

// InitAsyncPipelinePool Init one async pipeline pool
// It will do nothing if it's repeated init.
func InitAsyncPipelinePool(typ AsyncPipelineType) {
	if _, ok := asyncPipelinePools[typ]; ok {
		return
	}
	options := newAsyncPipelineOptions(typ)
	pipelines := make([]*asyncPipeline, 0, options.parallel)
	for i := 0; i < options.parallel; i++ {
		pipelines = append(pipelines, newAsyncPipeline(options))
	}

	asyncPipelinePools[typ] = &asyncPipelinePool{
		options:   options,
		pipelines: pipelines,
	}
}

// CloseAsyncPipeline exec all queued cmds and stop receiving any cmd
func CloseAsyncPipeline() {
	for _, pool := range asyncPipelinePools {
		for _, pipeline := range pool.pipelines {
			pipeline.Close()
		}
	}
	asyncPipelinePools = nil
}

func configStr(args ...string) string {
	return viper.GetString(strings.Join(args, "."))
}