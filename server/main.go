package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/cluster"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/modules"
	"github.com/topfreegames/pitaya/serialize/protobuf"
	"server/jcak_constants"
	"server/module_redis"
	"server/server_http"
	"server/server_lobby"
)

func configureLobby() {
	componentLobby := server_lobby.NewComponentLobby()
	pitaya.Register(componentLobby, component.WithName("ComponentLobby"), )
	pitaya.RegisterRemote(componentLobby, component.WithName("ComponentLobby"), )
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
	case jcak_constants.SvTypeGame:
		//configureGame()
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

	confs := viper.New()
	//confs.Set("pitaya.cluster.rpc.server.grpc.port", *port)

	meta := map[string]string{
		//constants.GRPCHostKey: "127.0.0.1",
		//constants.GRPCPortKey: strconv.Itoa(*port),
	}

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, meta, confs)
	gs, err := cluster.NewNatsRPCServer(pitaya.GetConfig(), pitaya.GetServer(), pitaya.GetMetricsReporters(), pitaya.GetDieChan())
	if err != nil {
		panic(err)
	}

	bs := modules.NewETCDBindingStorage(pitaya.GetServer(), pitaya.GetConfig())
	pitaya.RegisterModule(bs, "bindingsStorage")

	//注册redis
	rs := module_redis.NewRedisStorage(pitaya.GetConfig())
	pitaya.RegisterModule(rs, "redisStorage")

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
