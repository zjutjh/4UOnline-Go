package redis

import "4u-go/config/config"

func getConfig() redisConfig {
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
