package db_handler

import (
	"fmt"
	"github.com/topfreegames/pitaya/modules"
	"server/dao/redis_helper"
	"strconv"
)

type RedisRegisterHandler struct {
	modules.Base
	key string
}

func NewDBRegisterHandler() *RedisRegisterHandler {
	r := &RedisRegisterHandler{
		key:"db_handler_register",
	}
	return r
}

func (this *RedisRegisterHandler) accountKey(account string) string {
	return fmt.Sprintf("register[%v]", account)
}

// CheckIsRegister 检查是否已注册
func (this *RedisRegisterHandler) CheckIsRegister(account string) (bool, error) {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return false, err
	}
	exist, err := c.Exists(this.accountKey(account)).Result()
	if err != nil {
		return false, err
	}
	if exist != 0 {
		return true, nil
	}
	return false, nil
}

// GetRegisterInfoByAccount 通过account获取密码及用户Id
func (this *RedisRegisterHandler) GetRegisterInfoByAccount(account string) (password string, id int64, err error) {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return "", 0, err
	}
	data, err := c.HMGet(this.accountKey(account), "password", "id").Result()
	if err != nil {
		return "", 0, err
	}
	password = data[0].(string)
	pidStr := data[1].(string)
	id, err = strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		return "", 0, err
	}
	return password, id, nil
}

//NewRegister 注册新的账号
func (this *RedisRegisterHandler) NewRegister(account, password string, id int64) error {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return err
	}
	_, err = c.HMSet(this.accountKey(account), map[string]interface{}{
		"account":  account,
		"password": password,
		"id":       id,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}
