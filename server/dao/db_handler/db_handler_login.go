package db_handler

import (
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya/modules"
	"server/dao/redis_helper"
	"strconv"
)

type RedisLoginHandler struct {
	modules.Base

}

func NewDBLoginHandler() *RedisLoginHandler {
	return &RedisLoginHandler{

	}
}

func (this *RedisLoginHandler) loginIdKey() string {
	return "login_id_token"
}

func (this *RedisLoginHandler) loginTokenKey() string {
	return "login_token_id"
}

// IsAlreadyLogin 是否已登录
func (this *RedisLoginHandler) IsAlreadyLogin(token string) (bool, int64, error) {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err == nil {
		if uidStr, err := c.HGet(this.loginTokenKey(), token).Result(); err == nil {
			if uid, err := strconv.ParseInt(uidStr, 10, 64); err == nil {
				return true, uid, nil
			}
		}
	}
	return false, 0, err
}

//DeleteTokenByUID 如果存在老的token就删掉，如果没有也无所谓，所有如果是redis.Nil那么直接返回nil
func (this *RedisLoginHandler) DeleteTokenByUID(uid int64) error {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return err
	}
	uidStr := strconv.FormatInt(uid, 10)
	oldToken, err := c.HGet(this.loginIdKey(), uidStr).Result()
	if err != nil && err != redis.Nil {
		return err
	} else if err == redis.Nil {
		return nil
	} else {
		_, err := c.HDel(this.loginTokenKey(), oldToken).Result()
		if err != nil && err != redis.Nil {
			return err
		} else {
			return nil
		}
	}
}

//SaveLoginDB 保存登录信息
func (this *RedisLoginHandler) SaveLoginDB(uid int64, token string) error {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return err
	}
	pipe := c.Pipeline()
	uidStr := strconv.FormatInt(uid, 10)
	pipe.HSet(this.loginTokenKey(), token, uidStr)
	pipe.HSet(this.loginIdKey(), uidStr, token)
	_, err = pipe.Exec()
	if err != nil {
		return err
	}
	return nil
}


