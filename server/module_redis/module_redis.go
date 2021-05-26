package module_redis

import (
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/modules"
	"server/dao/redis_helper"
	"time"
)

type redisStorage struct {
	modules.Base
	addr           string
	db             int
	dialTimeout    time.Duration
	expirationTime time.Duration
	client         *redis.Client
}

func NewRedisStorage() *redisStorage {
	r := &redisStorage{
		addr:           pitaya.GetConfig().GetString("pitaya.modules.redisStorage.client.addr"),
		db:             pitaya.GetConfig().GetInt("pitaya.modules.redisStorage.client.db"),
		dialTimeout:    pitaya.GetConfig().GetDuration("pitaya.modules.redisStorage.client.dialTimeout"),
		expirationTime: pitaya.GetConfig().GetDuration("pitaya.modules.redisStorage.client.expirationTime"),
	}
	return r
}

func GetRedisStorage() *redisStorage {
	if targetModule, err := pitaya.GetModule("redisStorage"); err == nil {
		return targetModule.(*redisStorage)
	}
	return nil
}

func (this *redisStorage) Init() error {
	redis_helper.Init(redis_helper.Default)
	redis_helper.Init(redis_helper.User)
	return nil
}
//
//func (this *redisStorage) Add( data pojo2.DbData) error {
//	if data == nil || len(data.GetKey()) < 0 {
//		return errors.New("data is nil or empty!")
//	}
//	data.SetStatus(pojo2.OptInsert)
//	bytes, err := json.Marshal(data)
//	if err != nil {
//		logger.Log.Errorf("AddCache failed,error:%s", err.Error())
//		return err
//	}
//	this.client.Set(data.GetKey(), bytes, this.expirationTime)
//
//	return nil
//}
//
//func (this *redisStorage) Delete(data pojo2.DbData) error {
//	if data == nil || len(data.GetKey()) < 0 {
//		return errors.New("data is nil or empty!")
//	}
//	_, err := this.client.Del(data.GetKey()).Result()
//	if err != nil {
//		logger.Log.Errorf("DeleteCache failed,error:%s", err.Error())
//		return err
//	}
//	return nil
//}
//
//func (this *redisStorage) Update(data pojo2.DbData) error {
//	if data == nil || len(data.GetKey()) < 0 {
//		return errors.New("data is nil or empty!")
//	}
//	data.SetStatus(pojo2.OptUpdate)
//	bytes, err := json.Marshal(data)
//	if err != nil {
//		logger.Log.Errorf("update redis failed,error:%s", err.Error())
//		return err
//	}
//	_, err = this.client.Set(data.GetKey(), bytes, this.expirationTime).Result()
//	if err != nil {
//		logger.Log.Errorf("update redis failed,error:%s", err.Error())
//		return err
//	}
//	return nil
//}
//
//func (this *redisStorage) Find(key string, target interface{}) error {
//	result, err := this.client.Get(key).Bytes()
//	if err != nil {
//		logger.Log.Infof("redis result find error, key=:%s", key)
//		return err
//	}
//	if err := json.Unmarshal(result, target); err != nil {
//		logger.Log.Errorf("find result and json Unmarshal failed ", err.Error())
//		return err
//	}
//	return nil
//}
//
////
////func (this *redisStorage)IsHave(key string,db pojo.DbData)bool  {
////	c, err := this.client.(redis.Default)
////	if err != nil {
////		return false, err
////	}
////	exist, err := this.client.Exists(context.Background(),db.GetKey()).Result()
////	if err != nil {
////		return false, err
////	}
////	if exist != 0 {
////		return true, nil
////	} else {
////		return false, nil
////	}
////}