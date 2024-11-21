package lostAndFoundController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/lostAndFoundService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewLostAndFoundData struct {
	ID         uint `json:"id" binding:"required"`
	IsApproved bool `json:"is_approved"`
}

// ReviewLostAndFound 审核失物招领
func ReviewLostAndFound(c *gin.Context) {
	var data reviewLostAndFoundData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断失物招领是否存在
	_, err = lostAndFoundService.GetLostAndFoundById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	err = lostAndFoundService.ReviewLostAndFound(data.ID, data.IsApproved)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
