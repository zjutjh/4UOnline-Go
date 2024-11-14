package activityController

import (
	"errors"
	"time"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/activityService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateActivityData struct {
	ID           uint   `json:"id" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
	Department   string `json:"department" binding:"required"`
	StartTime    string `json:"start_time" binding:"required"`
	EndTime      string `json:"end_time" binding:"required"`
	Campus       uint8  `json:"campus" binding:"required"`
	Location     string `json:"location" binding:"required"`
	Img          string `json:"img"`
}

// UpdateActivity 更新校园活动
func UpdateActivity(c *gin.Context) {
	var data updateActivityData
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

	activity, err := activityService.GetActivityById(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	user := c.GetUint("user_id")
	adminType := c.GetUint("admin_type")
	if activity.AuthorID != user && adminType != models.SuperAdmin {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	{ // 更新活动信息
		activity.Title = data.Title
		activity.Introduction = data.Introduction
		activity.Department = data.Department
		activity.StartTime = startTime
		activity.EndTime = endTime
		activity.Campus = data.Campus
		activity.Location = data.Location
		activity.Img = data.Img
	}

	err = activityService.SaveActivity(activity)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
