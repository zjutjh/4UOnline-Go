package qrcodeController

import (
	"errors"
	"time"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/collegeService"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type getQrcodeData struct {
	ID uint `form:"id" binding:"required"`
}

type qrcodeResp struct {
	ID           uint           `json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	FeedbackType uint           `json:"feedback_type"` // 反馈类型
	College      models.College `json:"college"`       // 责任部门
	Department   string         `json:"department"`    // 负责单位
	Location     string         `json:"location"`      // 投放位置
	Status       bool           `json:"status"`        // 状态(是否启用)
	Description  string         `json:"description"`   // 备注

	ScanCount     uint `json:"scan_count"`     // 扫描次数
	FeedbackCount uint `json:"feedback_count"` // 反馈次数
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

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	resp, err := generateResp(qrcode)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, resp)
}

func generateResp(qrcode models.Qrcode) (*qrcodeResp, error) {
	college, err := collegeService.GetCollegeById(qrcode.College)
	if err != nil {
		return nil, err
	}

	return &qrcodeResp{
		ID:           qrcode.ID,
		CreatedAt:    qrcode.CreatedAt,
		FeedbackType: qrcode.FeedbackType,
		College:      college,
		Department:   qrcode.Department,
		Location:     qrcode.Location,
		ScanCount:    qrcode.ScanCount,
		Status:       qrcode.Status,

		FeedbackCount: qrcode.FeedbackCount,
		Description:   qrcode.Description,
	}, nil
}
