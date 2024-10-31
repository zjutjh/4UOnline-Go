package activityController

import (
	"strings"
	"time"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/activityService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createActivityData struct {
	Title        string   `json:"title" binding:"required"`
	Introduction string   `json:"introduction" binding:"required"`
	Department   string   `json:"department" binding:"required"`
	StartTime    string   `json:"start_time" binding:"required"`
	EndTime      string   `json:"end_time" binding:"required"`
	Campus       uint8    `json:"campus" binding:"required"`
	Location     string   `json:"location" binding:"required"`
	Photo        []string `json:"photo"`
}

// CreateActivity 创建一条校园活动
func CreateActivity(c *gin.Context) {
	var data createActivityData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ParamError, utils.LevelInfo, err)
		return
	}

	// 转换时间
	startTime, err := time.Parse(time.RFC3339, data.StartTime)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ParamError, utils.LevelInfo, err)
		return
	}
	endTime, err := time.Parse(time.RFC3339, data.EndTime)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ParamError, utils.LevelInfo, err)
		return
	}

	err = activityService.SaveActivity(models.Activity{
		Title:        data.Title,
		Introduction: data.Introduction,
		Department:   data.Department,
		StartTime:    startTime,
		EndTime:      endTime,
		Campus:       data.Campus,
		Location:     data.Location,
		Imgs:         strings.Join(data.Photo, ","),
		AuthorID:     c.GetUint("user_id"),
	})
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError, utils.LevelError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
