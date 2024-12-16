package announcementController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/announcementService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateAnnouncementData struct {
	ID         uint   `json:"id" binding:"required"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	Department string `json:"department" binding:"required"`
}

// UpdateAnnouncement 创建一条公告通知
func UpdateAnnouncement(c *gin.Context) {
	var data updateAnnouncementData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	announcement, err := announcementService.GetAnnouncementById(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	user := utils.GetUser(c)
	if announcement.AuthorID != user.ID && user.Type != models.SuperAdmin {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	{ // 更新公告信息
		announcement.Title = data.Title
		announcement.Content = data.Content
		announcement.Department = data.Department
	}

	err = announcementService.SaveAnnouncement(announcement)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
