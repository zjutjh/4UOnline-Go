package websiteController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/websiteService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type updateWebsiteData struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        uint   `json:"type" binding:"required"`
	College     uint   `json:"college"`
	Description string `json:"description" binding:"required"`
	Condition   string `json:"condition" binding:"required"`
	URL         string `json:"url" binding:"required"`
}

// UpdateWebsite 更新一个网站
func UpdateWebsite(c *gin.Context) {
	var data updateWebsiteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	website, err := websiteService.GetWebsiteById(data.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	user := utils.GetUser(c)
	if website.AuthorID != user.ID && user.Type != models.SuperAdmin {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	{ // 更新网站信息
		website.Title = data.Title
		website.Type = data.Type
		website.College = data.College
		website.Description = data.Description
		website.Condition = data.Condition
		website.URL = data.URL
	}

	err = websiteService.SaveWebsite(website)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
