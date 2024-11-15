package activityController

import (
	"time"

	"4u-go/app/apiException"
	"4u-go/app/services/activityService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getActivityResponse struct {
	ActivityList []activityElement `json:"activity_list"`
}

type activityElement struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Department string  `json:"department"`
	StartTime  string  `json:"start_time"`
	Campus     []uint8 `json:"campus"`
	Img        string  `json:"img"`
}

// GetActivityList 获取校园活动列表
func GetActivityList(c *gin.Context) {
	list, err := activityService.GetActivityList()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	activityList := make([]activityElement, 0)
	for _, activity := range list {
		activityList = append(activityList, activityElement{
			ID:         activity.ID,
			Title:      activity.Title,
			Department: activity.Department,
			StartTime:  activity.StartTime.Format(time.RFC3339),
			Campus:     activity.Campus,
			Img:        activity.Img,
		})
	}

	utils.JsonSuccessResponse(c, getActivityResponse{
		ActivityList: activityList,
	})
}
