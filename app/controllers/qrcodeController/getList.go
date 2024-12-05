package qrcodeController

import (
	"4u-go/app/apiException"
	"4u-go/app/services/qrcodeService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type filter struct {
	College      []uint `json:"college"`
	FeedbackType []uint `json:"feedback_type"`
}

type getListData struct {
	Keyword  string `json:"keyword"`
	Filter   filter `json:"filter"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type getListResponse struct {
	QrcodeList []qrcodeResp `json:"qrcode_list"`
	Total      int64        `json:"total"`
}

// GetList 实现了权益码列表的分页获取, 搜索, 筛选
func GetList(c *gin.Context) {
	var data getListData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	filter := data.Filter

	qrcodeListResp := make([]qrcodeResp, 0)

	qrcodeList, total, err := qrcodeService.GetList(
		filter.College, filter.FeedbackType,
		data.Keyword, data.Page, data.PageSize)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	for _, qrcode := range qrcodeList {
		resp, err := generateResp(qrcode)
		if err != nil {
			apiException.AbortWithException(c, apiException.ParamError, err)
			return
		}
		qrcodeListResp = append(qrcodeListResp, *resp)
	}

	utils.JsonSuccessResponse(c, getListResponse{
		QrcodeList: qrcodeListResp,
		Total:      total,
	})
}
