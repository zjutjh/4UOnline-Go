package userController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/userCenterService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

type changePasswordData struct {
	StudentID  string `json:"student_id" binding:"required"`
	IdentityID string `json:"identity_id" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// ChangePassword 修改密码接口
func ChangePassword(c *gin.Context) {
	var data changePasswordData
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

	err = userCenterService.RePassWithoutEmail(data.StudentID, data.IdentityID, data.Password)
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
