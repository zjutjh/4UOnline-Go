package redis

import "github.com/go-redis/redis/v8"

// redisConfig 定义 Redis 数据库的配置结构体
type redisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

// RedisClient 是全局的 Redis 客户端实例
var RedisClient *redis.Client

// RedisInfo 保存当前 Redis 配置信息
var RedisInfo redisConfig

// init 函数用于初始化 Redis 客户端和配置信息
func init() {
	info := getConfig()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     info.Host + ":" + info.Port,
		Password: info.Password,
		DB:       info.DB,
	})
	RedisInfo = info
}
