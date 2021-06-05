package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/constants"
	"github.com/topfreegames/pitaya/modules"
	"github.com/topfreegames/pitaya/serialize/protobuf"
	"server/dao/redis_module"
	"server/jcak_constants"
	"server/server_bjl"
	"server/server_http"
	"server/server_lobby"
)

func configureLobby() {
	componentLobby := server_lobby.NewComponentLobby()
	pitaya.Register(componentLobby, component.WithName("ComponentLobby"))
	pitaya.RegisterRemote(componentLobby, component.WithName("ComponentLobby"))
}

func configureBjl()  {
	componentLobby := server_bjl.NewComponentBjl()
	pitaya.Register(componentLobby, component.WithName("ComponentBjl"))
	pitaya.RegisterRemote(componentLobby, component.WithName("ComponentBjl"))
}

func configureHttpSever() {
	componentHttp := server_http.NewComponentHttp()
	pitaya.Register(componentHttp, component.WithName("ComponentHttp"))
	pitaya.RegisterRemote(componentHttp, component.WithName("ComponentHttp"))
}

func configureFrontend(port int) {
	tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", port))
	pitaya.AddAcceptor(tcp)
}

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")
	flag.Parse()
	defer pitaya.Shutdown()
	pitaya.SetSerializer(protobuf.NewSerializer())
	switch *svType {
	case jcak_constants.SvTypeConnector:
		configureFrontend(*port)
		break
	case jcak_constants.SvTypeLobby:
		configureLobby()
		break
	case jcak_constants.SvTypeBjl:
		configureBjl()
		break
	case jcak_constants.SvTypeWorld:
		//configureWorld()
		break
	case jcak_constants.SvTypeHttp:
		configureHttpSever()
		break
	default:
		fmt.Printf("error serverType = %s\n", *svType)
		return
	}
	configs:=setConfig()

	meta := map[string]string{
		constants.GRPCHostKey: "127.0.0.1", //正式项目中，需要设置grpc的host
		constants.GRPCPortKey: "3434",      //正式项目中，需要设置grpc的port
	}

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, meta, configs)

	registerModule()

	gs, err := cluster.NewNatsRPCServer(pitaya.GetConfig(), pitaya.GetServer(), pitaya.GetMetricsReporters(), pitaya.GetDieChan())
	if err != nil {
		panic(err)
	}

	gc, err := cluster.NewNatsRPCClient(
		pitaya.GetConfig(),
		pitaya.GetServer(),
		pitaya.GetMetricsReporters(),
		pitaya.GetDieChan(),
	)
	if err != nil {
		panic(err)
	}
	pitaya.SetRPCServer(gs)
	pitaya.SetRPCClient(gc)
	pitaya.Start()
}

func setConfig() *viper.Viper {
	//具体默认配置见：https://pitaya.readthedocs.io/en/latest/configuration.html
	configs := viper.New()

	//1.redis各项配置
	//1.1 redis的host
	configs.Set("pitaya.modules.redis.default.client.host", "127.0.0.1")
	//1.1 redis的port
	configs.Set("pitaya.modules.redis.default.client.port", 6379)
	//1.1 redis的哪个数据库，这里使用默认0的
	configs.Set("pitaya.modules.redis.default.client.db", 0)
	//1.1 redis的重试次数
	configs.Set("pitaya.modules.redis.default.client.retry",5)
	//1.1 redis的size,用于初始化redis
	configs.Set("pitaya.modules.redis.default.client.size",500)
	//1.1
	configs.Set("pitaya.modules.redis.default.client.idle",10)

	//2.grpc的配置
	//2.1.grpc的端口（实现上在默认的设置中，端口即是3434，见上面的配置地址）
	configs.Set("pitaya.cluster.rpc.server.grpc.port", 3434)
	return configs
}

func registerModule()  {
	//绑定etcd的模块
	bs := modules.NewETCDBindingStorage(pitaya.GetServer(), pitaya.GetConfig())
	pitaya.RegisterModule(bs, "bindingsStorage")

	//处理注册redis的模块
	redisModule := redis_module.NewRedisModule()
	pitaya.RegisterModule(redisModule,"redisModule")
}