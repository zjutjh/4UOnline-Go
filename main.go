package main

import (
	"log"

	"4u-go/app/midwares"
	"4u-go/app/utils/server"
	"4u-go/config/config"
	"4u-go/config/database"
	"4u-go/config/router"
	"4u-go/config/sdk"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 如果配置文件中开启了调试模式
	if !config.Config.GetBool("server.debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	if err := sdk.ZapInit(); err != nil {
		log.Fatal(err.Error())
	}
	if err := database.Init(); err != nil {
		zap.L().Fatal(err.Error())
	}
	if err := sdk.Init(r); err != nil {
		zap.L().Fatal(err.Error())
	}
	router.Init(r)

	server.Run(r, ":"+config.Config.GetString("server.port"))
}
