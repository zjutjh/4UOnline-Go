package qrcodeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type scanCountData struct {
	ID uint `form:"id" binding:"required"`
}

// ScanCount 更新权益码扫码次数
func ScanCount(c *gin.Context) {
	var data scanCountData
	if err := c.ShouldBind(&data); err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	if err := qrcodeService.AddScanCount(data.ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, nil)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
