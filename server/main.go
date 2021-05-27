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
	"server/dao/db_handler"
	"server/jcak_constants"
	"server/server_http"
	"server/server_lobby"
)

func configureLobby() {
	componentLobby := server_lobby.NewComponentLobby()
	pitaya.Register(componentLobby, component.WithName("ComponentLobby"))
	pitaya.RegisterRemote(componentLobby, component.WithName("ComponentLobby"))
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

	//redis conf
	configs.Set("pitaya.modules.redisStorage.client.host", "127.0.0.1")
	configs.Set("pitaya.modules.redisStorage.client.port", 6379)
	configs.Set("pitaya.modules.redisStorage.client.db", 0)
	configs.Set("pitaya.modules.redisStorage.client.retry",5)
	configs.Set("pitaya.modules.redisStorage.client.size",500)
	configs.Set("pitaya.modules.redisStorage.client.idle",10)
	configs.Set("pitaya.modules.redisStorage.client.idle","")

	//grpc的端口（实现上在默认的设置中，端口即是3434，见上面的配置地址）
	configs.Set("pitaya.cluster.rpc.server.grpc.port", 3434)
	return configs
}

func registerModule()  {

	//绑定etcd的模块
	bs := modules.NewETCDBindingStorage(pitaya.GetServer(), pitaya.GetConfig())
	pitaya.RegisterModule(bs, "bindingsStorage")

	//处理注册redis的模块
	dbHandlerRegister := db_handler.NewDBRegisterHandler()
	pitaya.RegisterModule(dbHandlerRegister,"dbHandlerRegister")

	//处理登录redis的模块
	dbHandlerLogin := db_handler.NewDBLoginHandler()
	pitaya.RegisterModule(dbHandlerLogin,"dbHandlerLogin")
}