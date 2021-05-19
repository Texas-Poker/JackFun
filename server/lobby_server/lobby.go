package lobby_server

import (
	"context"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/timer"
	"server/pb/pb_lobby"
)

type LobbyServer struct {
	component.Base
	timer *timer.Timer
}

func NewLobbyServer() *LobbyServer {
	return &LobbyServer{}
}



func (*LobbyServer) ReqLobbyInfo(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrorCode: "success"}, nil
}
