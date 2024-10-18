package config

import (
	"context"
	"errors"
	"time"

	"4u-go/app/models"
	"4u-go/config/database"
	"4u-go/config/redis"
	"gorm.io/gorm"
)

// 上下文用于 Redis 操作
var ctx = context.Background()

// getConfig 从 Redis 获取配置，如果不存在则从数据库中获取，并缓存到 Redis
func getConfig(key string) string {
	val, err := redis.RedisClient.Get(ctx, key).Result()
	if err == nil {
		return val
	}
	print(err)
	var config = &models.Config{}
	database.DB.Model(models.Config{}).Where(
		&models.Config{
			Key: key,
		}).First(&config)

	redis.RedisClient.Set(ctx, key, config.Value, 0)
	return config.Value
}

// setConfig 设置指定的配置项，如果不存在则创建新的配置。
func setConfig(key, value string) error {
	redis.RedisClient.Set(ctx, key, value, 0)
	var config models.Config
	result := database.DB.Where("`key` = ?", key).First(&config)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		config = models.Config{
			Key:        key,
			Value:      value,
			UpdateTime: time.Now(),
		}
		result = database.DB.Create(&config)
	} else {
		config.Value = value
		config.UpdateTime = time.Now()
		result = database.DB.Updates(&config)
	}
	return result.Error
}

// checkConfig 检查指定的配置项是否存在于 Redis 中。
func checkConfig(key string) bool {
	intCmd := redis.RedisClient.Exists(ctx, key)
	return intCmd.Val() == 1
}

func delConfig(key string) error {
	redis.RedisClient.Del(ctx, key)
	res := database.DB.Where(&models.Config{
		Key: key,
	}).Delete(models.Config{})
	return res.Error
}
