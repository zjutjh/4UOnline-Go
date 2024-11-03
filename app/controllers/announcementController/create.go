package announcementController

import (
	"time"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/announcementService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createAnnouncementData struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// CreateAnnouncement 创建一条公告通知
func CreateAnnouncement(c *gin.Context) {
	var data createAnnouncementData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = announcementService.SaveAnnouncement(models.Announcement{
		Title:       data.Title,
		Content:     data.Content,
		PublishTime: time.Now(),
		AuthorID:    c.GetUint("user_id"),
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
