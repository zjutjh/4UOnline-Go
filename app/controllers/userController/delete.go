package userController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/userService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type deleteAccountData struct {
	StudentID  string `json:"student_id" binding:"required"`
	IdentityID string `json:"identity_id" binding:"required"`
}

// DeleteAccount 注销账户
func DeleteAccount(c *gin.Context) {
	var data deleteAccountData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	user := utils.GetUser(c)
	if user.StudentID != data.StudentID {
		apiException.AbortWithException(c, apiException.NotPermission, err)
		return
	}

	// 若不是普通用户则提示不存在
	if user.Type != models.Undergraduate && user.Type != models.Postgraduate {
		apiException.AbortWithException(c, apiException.UserNotFound, nil)
		return
	}

	err = userService.DeleteAccount(user, data.IdentityID)
	if err != nil {
		var apiErr *apiException.Error
		if errors.As(err, &apiErr) {
			apiException.AbortWithException(c, apiErr, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
