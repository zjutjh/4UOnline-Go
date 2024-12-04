package qrcodeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type toggleStatusData struct {
	ID     uint `json:"id" binding:"required"`
	Status bool `json:"status"`
}

// ToggleStatus 更新学院信息
func ToggleStatus(c *gin.Context) {
	var data toggleStatusData
	err := c.ShouldBindJSON(&data)
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

	{ // 更新权益码状态
		qrcode.Status = data.Status
	}

	err = qrcodeService.SaveQrcode(qrcode)

	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
