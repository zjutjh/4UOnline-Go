package userService

import (
	"errors"

	"4u-go/app/apiException"
	"4u-go/app/models"
	"4u-go/app/services/userCenterService"
	"4u-go/config/database"
	"gorm.io/gorm"
)

// CreateStudentUser 创建学生用户
func CreateStudentUser(
	studentID, password, idCardNumber, email, name, college string,
	usertype uint,
) (*models.User, error) {
	_, err := GetUserByStudentID(studentID)
	if err == nil {
		return nil, apiException.UserAlreadyExisted
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	err = userCenterService.RegWithoutVerify(studentID, password, idCardNumber, email, usertype)
	if err != nil && !errors.Is(err, apiException.ReactiveError) {
		return nil, err
	}

	user := &models.User{
		Name:      name,
		College:   college,
		Type:      usertype,
		StudentID: studentID,
	}

	err = EncryptUserKeyInfo(user)
	if err != nil {
		return nil, err
	}
	res := database.DB.Create(&user)

	return user, res.Error
}

// CreateStudentUserWechat 创建学生用户(含微信)
func CreateStudentUserWechat(
	studentID string,
	password string,
	userType uint,
	idCardNumber string,
	email string,
	name string,
	college string,
	wechatOpenID string,
) (*models.User, error) {
	_, err := GetUserByWechatOpenID(wechatOpenID)
	if err == nil {
		return nil, apiException.OpenIDError
	}
	user, err := CreateStudentUser(studentID, password, idCardNumber, email, name, college, userType)
	if err != nil && !errors.Is(err, apiException.ReactiveError) {
		return nil, err
	}
	user.WechatOpenID = wechatOpenID
	database.DB.Save(user)
	return user, nil
}
