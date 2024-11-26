package redis

import (
	"4u-go/config/config"
	"github.com/go-redis/redis/v8"
)

// redisConfig 定义 Redis 数据库的配置结构体
type redisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

// GlobalClient 全局 Redis 客户端实例
var GlobalClient *redis.Client

// InfoConfig 保存 Redis 配置信息
var InfoConfig redisConfig

// Init 函数用于初始化 Redis 客户端和配置信息
func Init() {
	info := redisConfig{
		Host:     config.Config.GetString("redis.host"),
		Port:     config.Config.GetString("redis.port"),
		DB:       config.Config.GetInt("redis.db"),
		Password: config.Config.GetString("redis.pass"),
	}

	GlobalClient = redis.NewClient(&redis.Options{
		Addr:     info.Host + ":" + info.Port,
		Password: info.Password,
		DB:       info.DB,
	})
	InfoConfig = info
}
