package userController

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/services/sessionService"
	"4u-go/app/services/userService"
	"4u-go/app/utils"
	"4u-go/config/wechat"
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
		utils.JsonErrorResponse(c, apiException.ParamError)
		return
	}

	user, err := userService.GetUserByStudentID(postForm.StudentID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.JsonErrorResponse(c, apiException.UserNotFound)
		return
	}
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}

	if err := userService.AuthenticateUser(user, postForm.Password); err != nil {
		userService.HandleError(c, err)
		return
	}

	err = sessionService.SetUserSession(c, user)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"studentID":  user.StudentID,
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
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
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"studentID":  user.StudentID,
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}

// WeChatLogin 微信登录
func WeChatLogin(c *gin.Context) {
	var postForm autoLoginForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ParamError)
		return
	}

	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.OpenIDError)
		return
	}

	user, err := userService.GetUserByWechatOpenID(session.OpenID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.JsonErrorResponse(c, apiException.UserNotFound)
		return
	} else if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}

	err = sessionService.SetUserSession(c, user)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":         user.ID,
			"studentID":  user.StudentID,
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})
}
