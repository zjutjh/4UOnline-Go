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

type deleteWebsiteData struct {
	ID uint `json:"id" binding:"required"`
}

// DeleteWebsite 删除一个网站
func DeleteWebsite(c *gin.Context) {
	var data deleteWebsiteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	// 判断网站是否存在
	website, err := websiteService.GetWebsiteById(data.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			apiException.AbortWithException(c, apiException.ResourceNotFound, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	user := c.GetUint("user_id")
	adminType := c.GetUint("admin_type")
	if website.AuthorID != user && adminType != models.SuperAdmin {
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}

	err = websiteService.DeleteWebsiteById(data.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
