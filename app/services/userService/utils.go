package userService

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/config"
	"4u-go/app/models"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

// DecryptUserKeyInfo 解密用户信息
func DecryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.PhoneNum != "" {
		slt := utils.AesDecrypt(user.PhoneNum, key)
		user.PhoneNum = slt[0 : len(slt)-len(user.StudentID)]
	}
}

// EncryptUserKeyInfo 加密用户信息
func EncryptUserKeyInfo(user *models.User) {
	key := config.GetEncryptKey()
	if user.PhoneNum != "" {
		user.PhoneNum = utils.AesEncrypt(user.PhoneNum+user.StudentID, key)
	}
}

// HandleError 处理错误并返回相应的错误响应
func HandleError(c *gin.Context, err error) {
	var apiErr *apiException.Error
	if errors.As(err, &apiErr) {
		utils.JsonErrorResponse(c, apiErr)
	} else {
		utils.JsonErrorResponse(c, apiException.ServerError)
	}
}
