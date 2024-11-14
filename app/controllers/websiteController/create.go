package websiteController

import (
	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/websiteService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type createWebsiteData struct {
	Title       string `json:"title" binding:"required"`
	Type        uint   `json:"type" binding:"required"`
	College     uint   `json:"college"`
	Description string `json:"description" binding:"required"`
	Condition   string `json:"condition" binding:"required"`
	URL         string `json:"url" binding:"required"`
}

// CreateWebsite 新建一个网站
func CreateWebsite(c *gin.Context) {
	var data createWebsiteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// TODO: 限制学院管理员只能发自己学院的网站

	err = websiteService.SaveWebsite(models.Website{
		Title:       data.Title,
		Type:        data.Type,
		College:     data.College,
		Description: data.Description,
		Condition:   data.Condition,
		URL:         data.URL,
		AuthorID:    utils.GetUser(c).ID,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
