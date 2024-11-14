package activityController

import (
	"errors"

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
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断活动是否存在
	activity, err := activityService.GetActivityById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	user := c.GetUint("user_id")
	adminType := c.GetUint("admin_type")
	if activity.AuthorID != user && adminType != 4 {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	err = activityService.DeleteActivityById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
