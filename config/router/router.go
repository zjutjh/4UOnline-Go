package router

import (
	"4u-go/app/controllers/activityController"
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

			user.POST("/login/wechat", userController.WeChatLogin)
			user.POST("/login", userController.AuthByPassword)
			user.POST("/login/session", userController.AuthBySession)
		}

		activity := api.Group("/activity")
		{
			activity.GET("", activityController.GetActivityList)
			activity.POST("", midwares.CheckAdmin, activityController.CreateActivity)
			activity.PUT("", midwares.CheckAdmin, activityController.UpdateActivity)
			activity.DELETE("", midwares.CheckAdmin, activityController.DeleteActivity)
		}
	}
}
