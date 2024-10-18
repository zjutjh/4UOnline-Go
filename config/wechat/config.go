package wechat

import (
	"fmt"
	"strings"

	"4u-go/config/config"
)

type wechatConfig struct {
	Driver    string
	AppId     string
	AppSecret string
}

func getConfigs() (wechatConfig, error) {
	wc := wechatConfig{}
	if !config.Config.IsSet("wechat.appid") {
		return wc, fmt.Errorf("wechat.appid configError")
	}
	if !config.Config.IsSet("wechat.appsecret") {
		return wc, fmt.Errorf("wechat.appsecret configError")
	}
	wc.AppId = config.Config.GetString("wechat.appid")
	wc.AppSecret = config.Config.GetString("wechat.appsecret")

	wc.Driver = string(Memory)
	if config.Config.IsSet("wechat.driver") {
		wc.Driver = strings.ToLower(config.Config.GetString("wechat.driver"))
	}
	return wc, nil
}
