package activityController

import (
	"4u-go/app/apiException"
	"4u-go/app/services/activityService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deleteActivityData struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteActivity 删除一条校园活动
func DeleteActivity(c *gin.Context) {
	var data deleteActivityData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ParamError, utils.LevelInfo, err)
		return
	}

	// 判断活动是否存在
	activity, err := activityService.GetActivityById(data.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, apiException.ActivityNotFound, utils.LevelInfo, err)
		} else {
			utils.JsonErrorResponse(c, apiException.ServerError, utils.LevelError, err)
		}
		return
	}

	user := c.GetUint("user_id")
	adminType := c.GetUint("admin_type")
	if activity.AuthorID != user && adminType != 4 {
		utils.JsonErrorResponse(c, apiException.NotPermission, utils.LevelInfo, nil)
		return
	}

	err = activityService.DeleteActivityById(data.ID)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError, utils.LevelError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
