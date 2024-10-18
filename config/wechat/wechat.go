package wechat

import (
	"fmt"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
)

// driver 类型表示会话存储驱动的名称
type driver string

// 定义支持的驱动类型常量
const (
	Memory driver = "memory"
	Redis  driver = "redis"
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
	case string(Redis):
		wcCache = setRedis()
	case string(Memory):
		wcCache = cache.NewMemory()
	default:
		return fmt.Errorf("wechat configError")
	}

	cfg := &miniConfig.Config{
		AppID:     config.AppId,
		AppSecret: config.AppSecret,
		Cache:     wcCache,
	}

	MiniProgram = wc.GetMiniProgram(cfg)
	return nil
}
