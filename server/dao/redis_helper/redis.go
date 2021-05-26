package redis_helper

import (
	"context"
	"server/system"
)

const (
	// Default stands for redis.default in viper
	Default ClientType = "default"

	// Cli stands for redis.cli in viper
	Cli ClientType = "cli"

	// User stands for pipeline.user in viper
	User AsyncPipelineType = "user"
)

var background = context.Background()

// Init init all redis pool
func Init(typ string) {
	InitWatch()
	system.SeriousPanic(func() {
		switch typ {
		case Default:
			InitClientPool(Default)
		case Cli:
			InitClientPool(Cli)
		case User:
			InitAsyncPipelinePool(User)
		}
	})
}

// Close closes all redis pool
func Close() {
	CloseAsyncPipeline()
	CloseClientPool()
}
