package router

import (
	"4u-go/app/controllers/userController"
	"4u-go/app/midwares"
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre, midwares.CheckInit)
	{
		user := api.Group("/user")
		{
			user.POST("/create/student/wechat", userController.BindOrCreateStudentUserFromWechat)
			user.POST("/create/student", userController.CreateStudentUser)
		}
	}
}
