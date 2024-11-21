package activityController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/activityService"
	"4u-go/app/services/objectService"
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

	user := utils.GetUser(c)
	if activity.AuthorID != user.ID && user.Type != models.SuperAdmin {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	// 删除活动对应的图片
	objectKey, ok := objectService.GetObjectKeyFromUrl(activity.Img)
	if ok {
		err = objectService.DeleteObject(objectKey)
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}

	err = activityService.DeleteActivityById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
