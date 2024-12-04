package qrcodeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getQrcodeData struct {
	ID uint `form:"id" binding:"required"`
}

// GetQrcode 获取权益码信息
func GetQrcode(c *gin.Context) {
	var data getQrcodeData
	err := c.ShouldBind(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	qrcode, err := qrcodeService.GetQrcodeById(data.ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		return
	}

	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, qrcode)
}
