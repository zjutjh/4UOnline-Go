package lostAndFoundController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/lostAndFoundService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewLostAndFoundData struct {
	ID         uint `json:"id" binding:"required"`
	IsApproved bool `json:"is_approved"`
}

// ReviewLostAndFound 审核失物招领
func ReviewLostAndFound(c *gin.Context) {
	var data reviewLostAndFoundData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断失物招领是否存在
	_, err = lostAndFoundService.GetLostAndFoundById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	err = lostAndFoundService.ReviewLostAndFound(data.ID, data.IsApproved)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type updateLostAndFoundData struct {
	ID           uint   `json:"id" binding:"required"`
	Type         bool   `json:"type"`         // 1-失物 0-寻物
	Name         string `json:"name"`         // 物品名称
	Introduction string `json:"introduction"` // 物品介绍
	Campus       uint8  `json:"campus"`       // 校区 1-朝晖 2-屏峰 3-莫干山
	Kind         uint8  `json:"kind"`         // 物品种类 1其他2证件3箱包4首饰5现金6电子产品7钥匙
	Place        string `json:"place"`        // 丢失或拾得地点
	Time         string `json:"time"`         // 丢失或拾得时间
	Imgs         string `json:"imgs"`         // 物品图片，多个图片以逗号分隔
	PickupPlace  string `json:"pickup_place"` // 失物领取地点
	Contact      string `json:"contact"`      // 联系方式
}

// UpdateLostAndFound 修改失物招领
func UpdateLostAndFound(c *gin.Context) {
	var data updateLostAndFoundData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断失物招领是否存在
	record, err := lostAndFoundService.GetLostAndFoundById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	user := utils.GetUser(c)
	if user.Type != models.SuperAdmin && user.Type != models.ForU {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	{ // 更新失物招领信息
		record.Type = data.Type
		record.Name = data.Name
		record.Introduction = data.Introduction
		record.Campus = data.Campus
		record.Kind = data.Kind
		record.Place = data.Place
		record.Time = data.Time
		record.Imgs = data.Imgs
		record.Contact = data.Contact
		record.IsApproved = 2
		record.IsProcessed = 2
	}

	err = lostAndFoundService.SaveLostAndFound(record)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type updateLostAndFoundStatusData struct {
	ID uint `json:"id" binding:"required"`
}

// UpdateLostAndFoundStatus 用户设置失物招领为已完成
func UpdateLostAndFoundStatus(c *gin.Context) {
	var data updateLostAndFoundStatusData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断失物招领是否存在
	record, err := lostAndFoundService.GetLostAndFoundById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	user := utils.GetUser(c)
	if user.StudentID != record.Publisher {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	{ // 更新失物招领信息
		record.IsProcessed = 1
	}

	err = lostAndFoundService.SaveLostAndFound(record)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
