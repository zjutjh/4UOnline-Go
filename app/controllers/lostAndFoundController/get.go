package lostAndFoundController

import (
	"4u-go/app/apiException"
	"4u-go/app/services/lostAndFoundService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type getLostAndFoundListData struct {
	Type   bool  `json:"type"`                      // 1-失物 0-寻物
	Campus uint8 `json:"campus" binding:"required"` // 校区 1-朝晖 2-屏峰 3-莫干山
	Kind   uint8 `json:"kind"`                      // 物品种类 0全部1其他2证件3箱包4首饰5现金6电子产品7钥匙
}
type getLostAndFoundListResponse struct {
	LostAndFoundList []lostAndFoundElement `json:"list"`
}
type lostAndFoundElement struct {
	ID           uint   `json:"id"`
	Imgs         string `json:"imgs"`
	Name         string `json:"name"`
	Place        string `json:"place"`
	Time         string `json:"time"`
	Introduction string `json:"introduction"`
}

// GetLostAndFoundList 获取失物招领列表
func GetLostAndFoundList(c *gin.Context) {
	var data getLostAndFoundListData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	list, err := lostAndFoundService.GetLostAndFoundList(data.Type, data.Campus, data.Kind)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	lostAndFoundList := make([]lostAndFoundElement, 0)
	for _, record := range list {
		lostAndFoundList = append(lostAndFoundList, lostAndFoundElement{
			ID:           record.ID,
			Imgs:         record.Imgs,
			Name:         record.Name,
			Place:        record.Place,
			Time:         record.Time,
			Introduction: record.Introduction,
		})
	}

	utils.JsonSuccessResponse(c, getLostAndFoundListResponse{
		LostAndFoundList: lostAndFoundList,
	})
}
