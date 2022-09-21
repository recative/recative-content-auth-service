package cache

import (
	"encoding/json"
	"fmt"
	"github.com/recative/recative-backend/pkg/logger"
	"github.com/recative/recative-backend/utils/must"
	"go.uber.org/zap"
	"time"
)

type Cache interface {
	Store(key string, value any, expireSecond int) error
	Load(key string, dest any) (isExist bool, err error)
	Delete(key string) error
	RawClient() *redis.Client
}

type Config struct {
	RedisUri string `env:"REDIS_URI"`
}

func New(config Config) Cache {
	option, err := redis.ParseURL(config.RedisUri)
	if err != nil {
		logger.Panic(fmt.Sprintln("parse redis uri failed: " + err.Error()))
	}
	client := redis.NewClient(option)
	must.String(client.Ping().Result())

	return &cache{
		client,
	}
}

type cache struct {
	Redis *redis.Client
}

var _ Cache = &cache{}

func (cache *cache) Store(key string, value any, expireSecond int) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		logger.DPanic("Store MarshalToString failed", zap.Error(err))
	}
	dur := time.Second * time.Duration(expireSecond)
	return cache.Redis.Set(key, bytes, dur).Err()
}

func (cache *cache) Delete(key string) error {
	return cache.Redis.Del(key).Err()
}

func (cache *cache) Load(key string, dest any) (isExist bool, err error) {
	bytes, err := cache.Redis.Get(key).Bytes()
	// TODO: refactor to ,err
	if err != nil {
		return false, err
	}
	if err := json.Unmarshal(bytes, dest); err != nil {
		return false, err
	}
	return true, nil
}

func (cache *cache) RawClient() *redis.Client {
	return cache.Redis
}
