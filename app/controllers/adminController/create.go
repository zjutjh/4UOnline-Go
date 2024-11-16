package adminController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/adminService"
	"4u-go/app/utils"
	"4u-go/config/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createAdminByKeyData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Key      string `json:"key" binding:"required"`
}

// CreateAdminByKey 通过密钥创建普通管理员
func CreateAdminByKey(c *gin.Context) {
	var data createAdminByKeyData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	key := config.Config.GetString("admin.key")
	if data.Key != key {
		apiException.AbortWithException(c, apiException.AdminKeyError, err)
		return
	}

	_, err = adminService.GetUserByUsername(data.Username)
	if err == nil {
		apiException.AbortWithException(c, apiException.AdminAlreadyExisted, err)
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	_, err = adminService.CreateAdminUser(data.Username, data.Password, models.ForU, "")
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"username": data.Username,
		"password": data.Password,
	})
}
