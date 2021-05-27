package server_lobby

import (
	"context"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"server/dao/db_handler"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
)

type ComponentLobby struct {
	component.Base
	db_login_handler  *db_handler.RedisLoginHandler
}

func (this *ComponentLobby) Init() {
	if dbHandler, err := pitaya.GetModule("dbLoginRegister"); err == nil {
		this.db_login_handler = dbHandler.(*db_handler.RedisLoginHandler)
	}
}

func NewComponentLobby() *ComponentLobby {
	return &ComponentLobby{}
}



func (this *ComponentLobby) Test(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrCode: pb_enum.ErrorCode_AuthFailed}, nil
}