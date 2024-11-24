package wechat

import (
	"errors"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

// 定义支持的驱动类型常量
const (
	Memory = "memory"
	Redis  = "redis"
)

// MiniProgram 是一个指向小程序实例的指针
var MiniProgram *miniprogram.MiniProgram

// Init 初始化微信小程序配置。
func Init() error {
	config, err := getConfigs()
	if err != nil {
		return err
	}
	wc := wechat.NewWechat()
	var wcCache cache.Cache
	switch config.Driver {
	case Redis:
		wcCache = setRedis()
	case Memory:
		wcCache = cache.NewMemory()
	default:
		return errors.New("wechat config error")
	}

	cfg := &miniConfig.Config{
		AppID:     config.AppId,
		AppSecret: config.AppSecret,
		Cache:     wcCache,
	}

	MiniProgram = wc.GetMiniProgram(cfg)
	return nil
}
