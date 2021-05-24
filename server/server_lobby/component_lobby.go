package server_lobby

import (
	"context"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/timer"
	"log"
	"server/pb/pb_common"
	"server/pb/pb_lobby"
)

type ComponentLobby struct {
	component.Base
	timer *timer.Timer
}

func NewComponentLobby() *ComponentLobby {
	return &ComponentLobby{}
}

func (this *ComponentLobby) ReqLogin(ctx context.Context, req *pb_lobby.ReqLogin) (*pb_lobby.RespLogin, error) {
	log.Println("account=",req.Account)
	log.Println("password=",req.Password)
	return &pb_lobby.RespLogin{ErrCode: pb_common.ErrorCode_EntryError}, nil
}

func (this *ComponentLobby) ReqLobbyInfo(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrCode: pb_common.ErrorCode_OK}, nil
}
