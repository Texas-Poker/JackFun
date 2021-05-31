package redis_module

import (
	"github.com/topfreegames/pitaya/modules"
	"server/dao/redis_helper"
)

type RedisModule struct {
	modules.Base
}

func NewRedisModule() *RedisModule {
	redis_helper.Init(redis_helper.User)
	redis_helper.Init(redis_helper.Default)
	return &RedisModule{

	}
}