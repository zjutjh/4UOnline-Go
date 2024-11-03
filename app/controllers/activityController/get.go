package activityController

import (
	"time"

	"4u-go/app/apiException"
	"4u-go/app/services/activityService"
	"4u-go/app/services/sessionService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getActivityData struct {
	Campus uint8 `json:"campus" binding:"required"`
}

type getActivityResponse struct {
	ActivityList []activityElement `json:"activity_list"`
}

type activityElement struct {
	ID           uint   `json:"id"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Department   string `json:"department"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	PublishTime  string `json:"publish_time"`
	Campus       uint8  `json:"campus"`
	Location     string `json:"location"`
	Img          string `json:"img"`
	Editable     bool   `json:"editable"`
}

// GetActivityList 获取校园活动列表
func GetActivityList(c *gin.Context) {
	var data getActivityData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	user, err := sessionService.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	list, err := activityService.GetActivityList(data.Campus)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	activityList := make([]activityElement, 0)
	for _, activity := range list {
		activityList = append(activityList, activityElement{
			ID:           activity.ID,
			Title:        activity.Title,
			Introduction: activity.Introduction,
			Department:   activity.Department,
			StartTime:    activity.StartTime.Format(time.RFC3339),
			EndTime:      activity.EndTime.Format(time.RFC3339),
			PublishTime:  activity.PublishTime.Format(time.RFC3339),
			Campus:       activity.Campus,
			Location:     activity.Location,
			Img:          activity.Img,
			Editable:     activity.AuthorID == user.ID || user.Type == 4,
		})
	}

	utils.JsonSuccessResponse(c, getActivityResponse{
		ActivityList: activityList,
	})
}
