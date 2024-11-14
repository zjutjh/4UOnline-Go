package collageController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/collageService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deleteCollageData struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteCollage 删除学院
func DeleteCollage(c *gin.Context) {
	var data deleteCollageData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断学院是否存在
	_, err = collageService.GetCollageById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	err = collageService.DeleteCollageById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
