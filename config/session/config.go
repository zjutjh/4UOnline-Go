package session

import (
	"strings"

	"4u-go/config/config"
)

// 定义会话驱动类型常量
const (
	Memory = "memory"
	Redis  = "redis"
)

// 默认会话名称
var defaultName = "wejh-session"

// sessionConfig 存储会话的配置
type sessionConfig struct {
	Driver string
	Name   string
}

// getConfig 获取会话配置
func getConfig() sessionConfig {
	wc := sessionConfig{}
	wc.Driver = Memory
	if config.Config.IsSet("session.driver") {
		wc.Driver = strings.ToLower(config.Config.GetString("session.driver"))
	}

	wc.Name = defaultName
	if config.Config.IsSet("session.name") {
		wc.Name = strings.ToLower(config.Config.GetString("session.name"))
	}

	return wc
}

// getRedisConfig 获取 Redis 配置
func getRedisConfig() redisConfig {
	info := redisConfig{
		Host:     "localhost",
		Port:     "6379",
		DB:       0,
		Password: "",
	}
	if config.Config.IsSet("redis.host") {
		info.Host = config.Config.GetString("redis.host")
	}
	if config.Config.IsSet("redis.port") {
		info.Port = config.Config.GetString("redis.port")
	}
	if config.Config.IsSet("redis.db") {
		info.DB = config.Config.GetInt("redis.db")
	}
	if config.Config.IsSet("redis.pass") {
		info.Password = config.Config.GetString("redis.pass")
	}
	return info
}
