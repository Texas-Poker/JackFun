package pojo

import (
	"fmt"
	"github.com/go-redis/redis"
	"server/dao/hashtree"
	"server/dao/redis_helper"
	"server/pb/pb_enum"
	"strconv"
)

//User 执久层用户数据结构User(pojo_user)
type User struct {
	hashtree.Root
	Id       hashtree.Int64  `bson:"Id"`
	Account  hashtree.String `bson:"Account"`
	Password hashtree.String `bson:"Password"`
	NickName hashtree.String `bson:"NickName"`
	Icon     hashtree.String `bson:"Icon"`
	Age      hashtree.Int64  `bson:"Age"`
	Sex      pb_enum.Sex     `bson:"Sex"`
	Lv       hashtree.Int64  `bson:"Lv"`
	Gold     hashtree.Int64  `bson:"Gold"`
	Diamond  hashtree.Int64  `bson:"Diamond"`
	Token    hashtree.String `bson:"Token"`
}

// ManualSave 手动存redis
func (this *User) ManualSave() error {
	hash, err := this.Dump()
	if err != nil {
		return err
	}
	if len(hash) > 0 {
		p, err := redis_helper.GetAsyncPipeliner(redis_helper.User, this.Id.Get())
		if err != nil {
			return err
		}
		var cmds []redis.Cmder
		for key, value := range hash {
			cmds = append(cmds, redis.NewCmd("HSet", fmt.Sprintf("user[%v]", this.Id.Get()), key, value))
		}
		return p.DoCmds(cmds)
	}
	return nil
}

// CheckAccountExist 检查用户名是否已经存在
func CheckAccountExist(account string) (bool, error) {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return false, err
	}
	exist, err := c.Exists(fmt.Sprintf("account[%v]", account)).Result()
	if err != nil {
		return false, err
	}
	if exist != 0 {
		return true, nil
	}
	return false, nil
}

// QueryAccount queries account data
func QueryAccount(account string) (password string,id int64,err error) {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return "", 0, err
	}
	data, err := c.HMGet(fmt.Sprintf("account[%v]", account), "password", "id").Result()
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

func NewUserAndSaveRedis(account, password string, id int64) error {
	c, err := redis_helper.GetClient(redis_helper.Default)
	if err != nil {
		return err
	}
	_, err = c.HMSet(fmt.Sprintf("account[%v]", account), map[string]interface{}{
		"account":  account,
		"password": password,
		"id":       id,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}
