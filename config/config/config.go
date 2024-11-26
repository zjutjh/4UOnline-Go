package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config 是用于读取和管理配置的 viper 实例
var Config = viper.New()

// init 函数用于初始化配置，设置配置文件的名称、类型和路径，并加载配置文件。
func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("Config not find", err)
	}
}
