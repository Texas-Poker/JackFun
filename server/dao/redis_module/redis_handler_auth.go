package redis_module

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"server/dao/redis_helper"
	"server/system/coroutine"
)

func (this *RedisModule) UserKey(uid int64) string {
	return fmt.Sprintf("User[%d]", uid)
}

func (this *RedisModule) SetUserByUId(uid int64, data map[string]string) error {
	p, err := redis_helper.GetAsyncPipeliner(redis_helper.User, uid)
	if err != nil {
		return err
	}
	var cmds []redis.Cmder
	for key, value := range data {
		cmds = append(cmds, redis.NewCmd("HSet", this.UserKey(uid), key, value))
	}
	return p.DoCmds(cmds)
}


//GetUserByUId get user data when user login.
func (this *RedisModule) GetUserByUId(uid int64) (map[string]string, error) {
	p, err := redis_helper.GetAsyncPipeliner(redis_helper.User, uid)
	if err != nil {
		return nil, err
	}

	var hash map[string]string
	err = coroutine.Start(func(co coroutine.ID) error {
		if err := p.DoCmdWithCallback(
			redis.NewStringStringMapCmd("HGetAll", this.UserKey(uid)),
			func(cmder redis.Cmder) error {
				cmd := cmder.(*redis.StringStringMapCmd)
				data, err := cmd.Result()
				if err != nil {
					return err
				}

				if _, ok := coroutine.Resume(co, data); !ok {
					return errors.Errorf("coroutine [%s] resume failed", co)
				}
				return nil
			},
		); err != nil {
			return err
		}
		data := coroutine.Yield(co)
		hash = data[0].(map[string]string)
		return nil
	})

	return hash, err
}