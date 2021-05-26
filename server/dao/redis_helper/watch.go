package redis_helper

import (
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya"
)

var watchRetryTimes int

func InitWatch() {
	watchRetryTimes = pitaya.GetConfig().GetInt("pitaya.modules.redisStorage.client.retry")
	if watchRetryTimes == 0 {
		panic("error watchRetryTimes")
	}
}

// Watch 使用redis的乐观锁 如果txf里执行成功 返回nil  否则返回err
func Watch(client *redis.Client, txf func(*redis.Tx) error, key ...string) error {
	return client.Watch(txf, key...)
}

// WatchWithRetry 带重试的Watch
func WatchWithRetry(client *redis.Client, txf func(*redis.Tx) error, key ...string) error {
	count := int(0)
	var err error
	for count < watchRetryTimes {
		err = Watch(client, txf, key...)
		if err == nil {
			return nil
		}
		count++
	}
	return err
}
