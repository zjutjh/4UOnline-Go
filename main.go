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
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	log.ZapInit()
	if err := database.Init(); err != nil {
		log.Logger.Fatal(err.Error()) // 在 main 函数中处理错误并终止程序
	}
	if err := session.Init(r); err != nil {
		log.Logger.Fatal(err.Error())
	}
	if err := wechat.Init(); err != nil {
		log.Logger.Fatal(err.Error())
	}
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
}
