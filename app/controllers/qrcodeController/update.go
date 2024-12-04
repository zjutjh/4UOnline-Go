package qrcodeController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateQrcodeData struct {
	ID           uint   `json:"id" binding:"required"`
	College      uint   `json:"college" binding:"required"`
	Department   string `json:"department" binding:"required"`
	Description  string `json:"description" binding:"required"`
	FeedbackType uint   `json:"feedback_type" binding:"required"`
	Location     string `json:"location" binding:"required"`
}

// UpdateQrcode 更新学院信息
func UpdateQrcode(c *gin.Context) {
	var data updateQrcodeData
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

	{ // 更新权益码信息
		qrcode.College = data.College
		qrcode.Department = data.Department
		qrcode.Description = data.Description
		qrcode.FeedbackType = data.FeedbackType
		qrcode.Location = data.Location
	}

	err = qrcodeService.SaveQrcode(qrcode)

	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
