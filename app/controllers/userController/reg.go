package userController

import (
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
	Email        string `json:"email"  binding:"required"`
	Code         string `json:"code"  binding:"required"`
}

// BindOrCreateStudentUserFromWechat 微信创建学生用户
func BindOrCreateStudentUserFromWechat(c *gin.Context) {
	var postForm createStudentUserWechatForm
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
	postForm.StudentID = strings.ToUpper(postForm.StudentID)
	postForm.IDCardNumber = strings.ToUpper(postForm.IDCardNumber)
	user, err := userService.CreateStudentUserWechat(
		postForm.Password,
		postForm.StudentID,
		postForm.Type,
		postForm.IDCardNumber,
		postForm.Email,
		session.OpenID)
	if err != nil {
		userService.HandleError(c, err)
		return
	}

	err = sessionService.SetUserSession(c, user)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

type createStudentUserForm struct {
	StudentID    string `json:"studentID"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
	Type         uint   `json:"type"  binding:"required"` // 用户类型 1-本科生 2-研究生
	IDCardNumber string `json:"idCardNumber"  binding:"required"`
	Email        string `json:"email"  binding:"required"`
}

// CreateStudentUser H5创建学生用户
func CreateStudentUser(c *gin.Context) {
	var postForm createStudentUserForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ParamError)
		return
	}
	postForm.StudentID = strings.ToUpper(postForm.StudentID)
	postForm.IDCardNumber = strings.ToUpper(postForm.IDCardNumber)
	user, err := userService.CreateStudentUser(
		postForm.StudentID,
		postForm.Password,
		postForm.IDCardNumber,
		postForm.Email,
		postForm.Type)
	if err != nil {
		userService.HandleError(c, err)
		return
	}

	err = sessionService.SetUserSession(c, user)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.ServerError)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
