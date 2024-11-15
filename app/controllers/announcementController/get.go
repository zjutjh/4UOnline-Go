package announcementController

import (
	"errors"
	"time"

	"4u-go/app/apiException"
	"4u-go/app/services/announcementService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getAnnouncementData struct {
	ID uint `json:"id" binding:"required"`
}

type getAnnouncementResponse struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	PublishTime string `json:"publish_time"`
	Department  string `json:"department"`
}

type getAnnouncementListResponse struct {
	AnnouncementList []announcementElement `json:"announcement_list"`
}

type announcementElement struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	PublishTime string `json:"publish_time"`
}

// GetAnnouncementList 获取公告列表
func GetAnnouncementList(c *gin.Context) {
	list, err := announcementService.GetAnnouncementList()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	announcementList := make([]announcementElement, 0)
	for _, announcement := range list {
		announcementList = append(announcementList, announcementElement{
			ID:          announcement.ID,
			Title:       announcement.Title,
			PublishTime: announcement.CreatedAt.Format(time.RFC3339),
		})
	}

	utils.JsonSuccessResponse(c, getAnnouncementListResponse{
		AnnouncementList: announcementList,
	})
}

// GetAnnouncement 获取公告详情
func GetAnnouncement(c *gin.Context) {
	var data getAnnouncementData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	announcement, err := announcementService.GetAnnouncementById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	utils.JsonSuccessResponse(c, getAnnouncementResponse{
		Title:       announcement.Title,
		Content:     announcement.Content,
		PublishTime: announcement.CreatedAt.Format(time.RFC3339),
		Department:  announcement.Department,
	})
}
