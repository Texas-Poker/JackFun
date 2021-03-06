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
		constants.GRPCHostKey: "127.0.0.1", //??????????????????????????????grpc???host
		constants.GRPCPortKey: "3434",      //??????????????????????????????grpc???port
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
	//????????????????????????https://pitaya.readthedocs.io/en/latest/configuration.html
	configs := viper.New()

	//1.redis????????????
	//1.1 redis???host
	configs.Set("pitaya.modules.redis.default.client.host", "127.0.0.1")
	//1.1 redis???port
	configs.Set("pitaya.modules.redis.default.client.port", 6379)
	//1.1 redis???????????????????????????????????????0???
	configs.Set("pitaya.modules.redis.default.client.db", 0)
	//1.1 redis???????????????
	configs.Set("pitaya.modules.redis.default.client.retry",5)
	//1.1 redis???size,???????????????redis
	configs.Set("pitaya.modules.redis.default.client.size",500)
	//1.1
	configs.Set("pitaya.modules.redis.default.client.idle",10)

	//2.grpc?????????
	//2.1.grpc?????????????????????????????????????????????????????????3434??????????????????????????????
	configs.Set("pitaya.cluster.rpc.server.grpc.port", 3434)
	return configs
}

func registerModule()  {
	//??????etcd?????????
	bs := modules.NewETCDBindingStorage(pitaya.GetServer(), pitaya.GetConfig())
	pitaya.RegisterModule(bs, "bindingsStorage")

	//????????????redis?????????
	redisModule := redis_module.NewRedisModule()
	pitaya.RegisterModule(redisModule,"redisModule")
}