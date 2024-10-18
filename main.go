package main

import (
	"log"

	"4u-go/app/midwares"
	"4u-go/config/database"
	"4u-go/config/router"
	"4u-go/config/session"
	"4u-go/config/wechat"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatal(err) // 在 main 函数中处理错误并终止程序
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	if err := session.Init(r); err != nil {
		log.Fatal(err)
	}
	if err := wechat.Init(); err != nil {
		log.Fatal(err)
	}
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
