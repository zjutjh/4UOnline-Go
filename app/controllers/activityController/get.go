package activityController

import (
	"errors"
	"time"

	"4u-go/app/apiException"
	"4u-go/app/services/activityService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getActivityListResponse struct {
	ActivityList []activityElement `json:"activity_list"`
}

type activityElement struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Department string `json:"department"`
	StartTime  string `json:"start_time"`
	Campus     []uint `json:"campus"`
	Img        string `json:"img"`
}

type getActivityData struct {
	ID uint `json:"id" binding:"required"`
}

type getActivityResponse struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Department   string `json:"department"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Campus       []uint `json:"campus"`
	Location     string `json:"location"`
	Img          string `json:"img"`
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
			Campus:     utils.DecodeCampus(activity.Campus),
			Img:        activity.Img,
		})
	}

	utils.JsonSuccessResponse(c, getActivityListResponse{
		ActivityList: activityList,
	})
}

// GetActivity 获取活动详情
func GetActivity(c *gin.Context) {
	var data getActivityData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	activity, err := activityService.GetActivityById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	utils.JsonSuccessResponse(c, getActivityResponse{
		Title:        activity.Title,
		Introduction: activity.Introduction,
		Department:   activity.Department,
		StartTime:    activity.StartTime.Format(time.RFC3339),
		EndTime:      activity.EndTime.Format(time.RFC3339),
		Campus:       utils.DecodeCampus(activity.Campus),
		Location:     activity.Location,
		Img:          activity.Img,
	})
}
