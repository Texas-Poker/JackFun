package server_lobby

import (
	"context"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"server/dao/redis_module"
	"server/pb/pb_enum"
	"server/pb/pb_lobby"
)

type ComponentLobby struct {
	component.Base
	redisModule *redis_module.RedisModule
}

func (this *ComponentLobby) Init() {
	if dbHandler, err := pitaya.GetModule("redisModule"); err == nil {
		this.redisModule = dbHandler.(*redis_module.RedisModule)
	}
}

func NewComponentLobby() *ComponentLobby {
	return &ComponentLobby{}
}



func (this *ComponentLobby) Test(ctx context.Context, req *pb_lobby.ReqLobbyInfo) (*pb_lobby.RespLobbyInfo, error) {
	return &pb_lobby.RespLobbyInfo{ErrCode: pb_enum.ErrorCode_AuthFailed}, nil
}