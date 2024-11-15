package activityController

import (
	"time"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/activityService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createActivityData struct {
	Title        string `json:"title" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Department   string `json:"department" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"`
	EndTime      string `json:"end_time" binding:"required"`
	Campus       []uint `json:"campus" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Img          string `json:"img"`
}

// CreateActivity 创建一条校园活动
func CreateActivity(c *gin.Context) {
	var data createActivityData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 转换时间
	startTime, err := time.Parse(time.RFC3339, data.StartTime)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	endTime, err := time.Parse(time.RFC3339, data.EndTime)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = activityService.SaveActivity(models.Activity{
		Title:        data.Title,
		Introduction: data.Introduction,
		Department:   data.Department,
		StartTime:    startTime,
		EndTime:      endTime,
		Campus:       utils.EncodeCampus(data.Campus),
		Location:     data.Location,
		Img:          data.Img,
		AuthorID:     utils.GetUser(c).ID,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
