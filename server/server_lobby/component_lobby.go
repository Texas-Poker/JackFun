package server_lobby

import (
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"server/dao/redis_module"
)

type ComponentLobby struct {
	component.Base
	redisModule *redis_module.RedisModule
}

func (this *ComponentLobby) Init() {
	if dbHandler, err := pitaya.GetModule("redisModule"); err == nil {
		this.redisModule = dbHandler.(*redis_module.RedisModule)
	}
	//scheduler.NewTimer()
}

func NewComponentLobby() *ComponentLobby {
	return &ComponentLobby{}
}



