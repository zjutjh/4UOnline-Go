package userController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/sessionService"
	"4u-go/app/services/userService"
	"4u-go/app/utils"
	"4u-go/config/sdk"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type passwordLoginForm struct {
	StudentID string `json:"student_id" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

// AuthByPassword 通过密码认证
func AuthByPassword(c *gin.Context) {
	var postForm passwordLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	user, err := userService.GetUserByStudentID(postForm.StudentID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.UserNotFound, err)
		return
	}
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	if err := userService.AuthenticateUser(user, postForm.Password); err != nil {
		var apiErr *apiException.Error
		if errors.As(err, &apiErr) {
			apiException.AbortWithException(c, apiErr, err)
		} else {
			apiException.AbortWithException(c, apiException.ServerError, err)
		}
		return
	}

	err = sessionService.SetUserSession(c, user)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"studentID": user.StudentID,
			"userType":  user.Type,
			"college":   user.College,
		},
	})
}

type autoLoginForm struct {
	Code string `json:"code" binding:"required"`
}

// AuthBySession 通过session认证
func AuthBySession(c *gin.Context) {
	user, err := sessionService.UpdateUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"studentID": user.StudentID,
			"userType":  user.Type,
			"college":   user.College,
		},
	})
}

// WeChatLogin 微信登录
func WeChatLogin(c *gin.Context) {
	var postForm autoLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	session, err := sdk.MiniProgram.GetAuth().Code2Session(postForm.Code)
	if err != nil {
		apiException.AbortWithException(c, apiException.OpenIDError, err)
		return
	}

	user, err := userService.GetUserByWechatOpenID(session.OpenID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.UserNotFound, err)
		return
	} else if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	err = sessionService.SetUserSession(c, user)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"studentID": user.StudentID,
			"userType":  user.Type,
			"college":   user.College,
		},
	})
}
