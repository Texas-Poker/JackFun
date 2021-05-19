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
	"server/constants"
	"server/lobby_server"
	"server/redis"
	"strings"
)


func configureLobby() {
	lobbyServer := lobby_server.NewLobbyServer()
	pitaya.Register(lobbyServer,
		component.WithName("lobby_server"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.RegisterRemote(lobbyServer,
		component.WithName("lobby_server"),
		component.WithNameFunc(strings.ToLower),
	)
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
	case constants.SvTypeConnector:
		configureFrontend(*port)
		break
	case constants.SvTypeLobby:
		configureLobby()
		break
	case constants.SvTypeGame:
		//configureGame()
		break
	case constants.SvTypeWorld:
		//configureWorld()
		break
	default:
		fmt.Printf("error svType %s\n", svType)
		return
	}

	confs := viper.New()


	meta := map[string]string{
		//constants.GRPCHostKey: "127.0.0.1",
		//constants.GRPCPortKey: *rpcServerPort,
	}

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, meta, confs)
	gs, err := cluster.NewNatsRPCServer(pitaya.GetConfig(), pitaya.GetServer(), pitaya.GetMetricsReporters(), pitaya.GetDieChan())
	if err != nil {
		panic(err)
	}

	bs := modules.NewETCDBindingStorage(pitaya.GetServer(), pitaya.GetConfig())
	pitaya.RegisterModule(bs, "bindingsStorage")

	//注册redis
	rs := redis.NewRedisStorage(pitaya.GetConfig())
	pitaya.RegisterModule(rs,"redisStorage")


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
