package wechat

import (
	"context"

	"4u-go/config/config"
	"4u-go/config/redis"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

// MiniProgram 是一个指向小程序实例的指针
var MiniProgram *miniprogram.MiniProgram

// Init 初始化微信小程序配置。
func Init() {
	info := redis.InfoConfig
	appId := config.Config.GetString("wechat.appid")
	appSecret := config.Config.GetString("wechat.appsecret")

	wc := wechat.NewWechat()
	wcCache := cache.NewRedis(context.Background(), &cache.RedisOpts{
		Host:        info.Host + ":" + info.Port,
		Database:    info.DB,
		MaxActive:   10,
		MaxIdle:     10,
		IdleTimeout: 60,
	})

	cfg := &miniConfig.Config{
		AppID:     appId,
		AppSecret: appSecret,
		Cache:     wcCache,
	}

	MiniProgram = wc.GetMiniProgram(cfg)
}
