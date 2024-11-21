package lostAndFoundController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/lostAndFoundService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

type getLostAndFoundContentData struct {
	ID uint `json:"id" binding:"required"`
}

// GetLostAndFoundContact 获取失物招领联系方式
func GetLostAndFoundContact(c *gin.Context) {
	var data getLostAndFoundContentData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	contact, err := lostAndFoundService.GetLostAndFoundContact(data.ID, utils.GetUser(c).StudentID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	utils.JsonSuccessResponse(c, contact)
}

type latestLostAndFoundResponse struct {
	Type         bool   `json:"type"`
	Imgs         string `json:"imgs"`
	Name         string `json:"name"`
	Place        string `json:"place"`
	Introduction string `json:"introduction"`
}

// GetLatestLostAndFound 获取最新失物招领
func GetLatestLostAndFound(c *gin.Context) {
	record, err := lostAndFoundService.GetLatestLostAndFound()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, latestLostAndFoundResponse{
		Type:         record.Type,
		Imgs:         record.Imgs,
		Name:         record.Name,
		Place:        record.Place,
		Introduction: record.Introduction,
	})
}

type getLostAndFoundStatusData struct {
	IsProcessed uint8 `json:"is_processed"` // 是否已处理 0-已取消 1-已处理 2-待处理
}
type getLostAndFoundStatusResponse struct {
	List []lostAndFoundStatusElement `json:"list"`
}
type lostAndFoundStatusElement struct {
	ID           uint   `json:"id"`
	Type         bool   `json:"type"`
	Imgs         string `json:"imgs"`
	Name         string `json:"name"`
	Kind         uint8  `json:"kind"`
	Place        string `json:"place"`
	Time         string `json:"time"`
	Introduction string `json:"introduction"`
}

// GetUserLostAndFoundStatus 查看发布失物招领信息后的审核状态
func GetUserLostAndFoundStatus(c *gin.Context) {
	var data getLostAndFoundStatusData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	list, err := lostAndFoundService.GetUserLostAndFoundStatus(utils.GetUser(c).StudentID, data.IsProcessed)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	lostAndFoundList := make([]lostAndFoundStatusElement, 0)
	for _, record := range list {
		lostAndFoundList = append(lostAndFoundList, lostAndFoundStatusElement{
			ID:           record.ID,
			Type:         record.Type,
			Imgs:         record.Imgs,
			Name:         record.Name,
			Kind:         record.Kind,
			Place:        record.Place,
			Time:         record.Time,
			Introduction: record.Introduction,
		})
	}

	utils.JsonSuccessResponse(c, getLostAndFoundStatusResponse{
		List: lostAndFoundList,
	})
}
