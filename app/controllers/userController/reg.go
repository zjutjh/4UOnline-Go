package userController

import (
	"errors"
	"strings"

	"4u-go/app/apiException"
	"4u-go/app/services/sessionService"
	"4u-go/app/services/userService"
	"4u-go/app/utils"
	"4u-go/config/wechat"
	"github.com/gin-gonic/gin"
)

type createStudentUserWechatForm struct {
	StudentID    string `json:"studentID"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	Type         uint   `json:"type"  binding:"required"` // 用户类型 1-本科生 2-研究生
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Name         string `json:"name"  binding:"required"`
	College      string `json:"college"  binding:"required"`
	Code         string `json:"code"  binding:"required"`
	Email        string `json:"email"`
}

// BindOrCreateStudentUserFromWechat 微信创建学生用户
func BindOrCreateStudentUserFromWechat(c *gin.Context) {
	var postForm createStudentUserWechatForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	session, err := wechat.MiniProgram.GetAuth().Code2Session(postForm.Code)
	if err != nil {
		apiException.AbortWithException(c, apiException.OpenIDError, err)
		return
	}
	postForm.StudentID = strings.ToUpper(postForm.StudentID)
	postForm.IDCardNumber = strings.ToUpper(postForm.IDCardNumber)
	user, err := userService.CreateStudentUserWechat(
		postForm.Password,
		postForm.StudentID,
		postForm.Type,
		postForm.IDCardNumber,
		postForm.Email,
		postForm.Name,
		postForm.College,
		session.OpenID)
	if err != nil {
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
	utils.JsonSuccessResponse(c, nil)
}

type createStudentUserForm struct {
	StudentID    string `json:"studentID"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	Type         uint   `json:"type"  binding:"required"` // 用户类型 1-本科生 2-研究生
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Name         string `json:"name"  binding:"required"`
	College      string `json:"college"  binding:"required"`
	Email        string `json:"email"`
}

// CreateStudentUser H5创建学生用户
func CreateStudentUser(c *gin.Context) {
	var postForm createStudentUserForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	postForm.StudentID = strings.ToUpper(postForm.StudentID)
	postForm.IDCardNumber = strings.ToUpper(postForm.IDCardNumber)
	user, err := userService.CreateStudentUser(
		postForm.StudentID,
		postForm.Password,
		postForm.IDCardNumber,
		postForm.Email,
		postForm.Name,
		postForm.College,
		postForm.Type,
	)
	if err != nil {
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
	utils.JsonSuccessResponse(c, nil)
}
