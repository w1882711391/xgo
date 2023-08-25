package dao

import (
    "fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func RedisInit() {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", opt.Cfg.Redis.Host, opt.Cfg.Redis.Port),
		Password: opt.Cfg.Redis.Password,
		DB:       0,
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic("Redis连接失败")
	}
}
