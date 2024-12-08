package qrcodeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deleteQrcodeData struct {
	ID uint `json:"id"`
}

// DeleteQrcode 删除一个权益码
func DeleteQrcode(c *gin.Context) {
	var data deleteQrcodeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断权益码是否存在
	_, err = qrcodeService.GetQrcodeById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	// 删除权益码
	err = qrcodeService.DeleteQrcodeById(data.ID)

	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
