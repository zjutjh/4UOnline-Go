package qrcodeController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createQrcodeData struct {
	College      uint   `json:"college" binding:"required"`
	Department   string `json:"department" binding:"required"`
	Description  string `json:"description" binding:"required"`
	FeedbackType uint   `json:"feedback_type" binding:"required"`
	Location     string `json:"location" binding:"required"`
}

// CreateQrcode 创建一个权益码
func CreateQrcode(c *gin.Context) {
	var data createQrcodeData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	err = qrcodeService.SaveQrcode(models.Qrcode{
		Status:       true,
		College:      data.College,
		Department:   data.Department,
		Description:  data.Description,
		FeedbackType: data.FeedbackType,
		Location:     data.Location,
	})

	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
