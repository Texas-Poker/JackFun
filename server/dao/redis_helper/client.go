package redis_helper

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/topfreegames/pitaya"
)

type (
	// ClientType is the type of client instance
	ClientType = string

	clientPool struct {
		options *redis.Options
		client  *redis.Client
	}
)

var (
	clientPools = make(map[string]*clientPool)
)

// GetClient is to get a client from redis pool
// Funcs of redis.Client are thread-safe
func GetClient(typ ClientType) (*redis.Client, error) {
	pool, err := getClientPool(typ)
	if err != nil {
		return nil, err
	}
	return pool.client, nil
}

func getClientPool(typ ClientType) (*clientPool, error) {
	pool, ok := clientPools[typ]
	if ok {
		return pool, nil
	}

	InitClientPool(typ)

	pool, ok = clientPools[typ]
	if !ok {
		return nil, fmt.Errorf("cannot find redis client type %s", typ)
	}
	return pool, nil
}

// NewClientOptions create new client options by client type
func NewClientOptions(typ ClientType) *redis.Options {
	config:=pitaya.GetConfig()
	configStr("pipeline", typ, "redis")
	host := config.GetString("pitaya.modules.redis.default.client.host")
	port := config.GetInt("pitaya.modules.redis.default.client.port")
	auth := config.GetString("pitaya.modules.redis.default.client.aut")
	db := config.GetInt("pitaya.modules.redis.default.client.db")
	size := config.GetInt("pitaya.modules.redis.default.client.size")
	idle := config.GetInt("pitaya.modules.redis.default.client.idle")

	return &redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     auth,
		DB:           int(db),
		PoolSize:     int(size),
		MinIdleConns: int(idle),
	}
}

// InitClientPool Init one redis pool
// It will do nothing if it's repeated init.
func InitClientPool(typ ClientType) {
	if _, ok := clientPools[typ]; ok {
		return
	}
	options := NewClientOptions(typ)
	client := redis.NewClient(options)
	clientPools[typ] = &clientPool{
		options: options,
		client:  client,
	}
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
}

// CloseClientPool closes all redis client pool
func CloseClientPool() {
	for _, pool := range clientPools {
		pool.client.Close()
	}
	clientPools = nil
}
