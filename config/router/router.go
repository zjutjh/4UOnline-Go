package router

import (
	"4u-go/app/controllers/activityController"
	"4u-go/app/controllers/adminController"
	"4u-go/app/controllers/announcementController"
	"4u-go/app/controllers/collegeController"
	"4u-go/app/controllers/userController"
	"4u-go/app/controllers/websiteController"
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

		admin := api.Group("/admin")
		{
			admin.POST("/create/key", adminController.CreateAdminByKey)
		}

		activity := api.Group("/activity")
		{
			activity.GET("", midwares.CheckLogin, activityController.GetActivityList)
			activity.POST("", midwares.CheckAdmin, activityController.CreateActivity)
			activity.PUT("", midwares.CheckAdmin, activityController.UpdateActivity)
			activity.DELETE("", midwares.CheckAdmin, activityController.DeleteActivity)
		}

		announcement := api.Group("/announcement")
		{
			announcement.GET("", midwares.CheckLogin, announcementController.GetAnnouncementList)
			announcement.POST("", midwares.CheckAdmin, announcementController.CreateAnnouncement)
			announcement.PUT("", midwares.CheckAdmin, announcementController.UpdateAnnouncement)
			announcement.DELETE("", midwares.CheckAdmin, announcementController.DeleteAnnouncement)
		}

		college := api.Group("/college")
		{
			college.GET("", collegeController.GetCollegeList)
			college.POST("", midwares.CheckSuperAdmin, collegeController.CreateCollege)
			college.PUT("", midwares.CheckSuperAdmin, collegeController.UpdateCollege)
			college.DELETE("", midwares.CheckSuperAdmin, collegeController.DeleteCollege)
		}

		website := api.Group("/website")
		{
			website.GET("", websiteController.GetWebsiteList)
			website.POST("", midwares.CheckAdmin, websiteController.CreateWebsite)
			website.DELETE("", midwares.CheckAdmin, websiteController.DeleteWebsite)
			website.PUT("", midwares.CheckAdmin, websiteController.UpdateWebsite)

			website.GET("/admin", midwares.CheckAdmin, websiteController.GetEditableWebsites)
		}
	}
}
