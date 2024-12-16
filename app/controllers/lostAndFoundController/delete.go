package lostAndFoundController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/lostAndFoundService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deleteLostAndFoundData struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteLostAndFound 撤回一条失物招领
func DeleteLostAndFound(c *gin.Context) {
	var data deleteLostAndFoundData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断失物招领是否存在
	record, err := lostAndFoundService.GetLostAndFoundById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	user := utils.GetUser(c)
	if user.Type != models.SuperAdmin && user.Type != models.ForU && user.StudentID != record.Publisher {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	err = lostAndFoundService.DeleteLostAndFoundById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
