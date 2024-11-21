package userController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/userCenterService"
	"4u-go/app/services/userService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type changePasswordData struct {
	StudentId  string `json:"student_id" binding:"required"`
	IdentityId string `json:"identity_id" binding:"required"`
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

	user, err := userService.GetUserByStudentID(data.StudentId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.UserNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	// 若不是普通用户则提示不存在
	if user.Type != models.Undergraduate && user.Type != models.Postgraduate {
		apiException.AbortWithException(c, apiException.UserNotFound, nil)
		return
	}

	err = userCenterService.RePassWithoutEmail(data.StudentId, data.IdentityId, data.Password)
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
