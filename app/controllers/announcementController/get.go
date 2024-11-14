package announcementController

import (
	"time"

	"4u-go/app/apiException"
	"4u-go/app/services/announcementService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getAnnouncementResponse struct {
	AnnouncementList []announcementElement `json:"announcement_list"`
}

type announcementElement struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	PublishTime string `json:"publish_time"`
	Editable    bool   `json:"editable"`
}

// GetAnnouncementList 获取公告列表
func GetAnnouncementList(c *gin.Context) {
	user := utils.GetUser(c)

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
			Content:     announcement.Content,
			PublishTime: announcement.PublishTime.Format(time.RFC3339),
			Editable:    announcement.AuthorID == user.ID || user.Type == 4,
		})
	}

	utils.JsonSuccessResponse(c, getAnnouncementResponse{
		AnnouncementList: announcementList,
	})
}
