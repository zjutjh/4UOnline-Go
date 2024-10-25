package main

import (
	"4u-go/app/midwares"
	"4u-go/app/utils/log"
	"4u-go/config/database"
	"4u-go/config/router"
	"4u-go/config/session"
	"4u-go/config/wechat"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	log.ZapInit()
	if err := database.Init(); err != nil {
		zap.L().Fatal(err.Error()) // 在 main 函数中处理错误并终止程序
	}
	if err := session.Init(r); err != nil {
		zap.L().Fatal(err.Error())
	}
	if err := wechat.Init(); err != nil {
		zap.L().Fatal(err.Error())
	}
	router.Init(r)

	err := r.Run()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
}
