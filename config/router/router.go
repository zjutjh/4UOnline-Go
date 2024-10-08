package router

import (
	"4u-go/app/midwares"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	const pre = "/api"

	api := r.Group(pre, midwares.CheckInit)
	{
		api.GET("")
	}
}
