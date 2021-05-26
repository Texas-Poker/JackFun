package server_lobby

import (
	"context"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/timer"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
)

type ComponentLobby struct {
	component.Base
	timer *timer.Timer
}

func NewComponentLobby() *ComponentLobby {
	return &ComponentLobby{}
}

func (this *ComponentLobby) ReqLobbyInfo(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrCode: pb_enum.ErrorCode_OK}, nil
}

func (this *ComponentLobby) Test(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrCode: pb_enum.ErrorCode_AuthFailed}, nil
}