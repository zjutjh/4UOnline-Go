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

type deleteAnnouncementData struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteAnnouncement 删除一条公告通知
func DeleteAnnouncement(c *gin.Context) {
	var data deleteAnnouncementData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断公告是否存在
	announcement, err := announcementService.GetAnnouncementById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	user := utils.GetUser(c)
	if announcement.AuthorID != user.ID && user.Type != models.SuperAdmin {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	err = announcementService.DeleteAnnouncementById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
