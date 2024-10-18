package wechat

import (
	"context"

	"4u-go/config/redis"
	"github.com/silenceper/wechat/v2/cache"
)

func setRedis() cache.Cache {
	redisOpts := &cache.RedisOpts{
		Host:        redis.RedisInfo.Host + ":" + redis.RedisInfo.Port,
		Database:    redis.RedisInfo.DB,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	}
	return cache.NewRedis(context.Background(), redisOpts)
}
