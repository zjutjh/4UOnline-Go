package userCenterService

import (
	"net/url"

	"4u-go/app/apiException"
	"4u-go/config/api/userCenterApi"
)

// Login 用户中心登录
func Login(stuId string, pass string) error {
	loginUrl, err := url.Parse(userCenterApi.Auth)
	if err != nil {
		return err
	}
	urlPath := loginUrl.String()
	regMap := map[string]any{
		"stu_id":       stuId,
		"password":     pass,
		"bound_system": 1,
	}
	resp, err := FetchHandleOfPost(regMap, urlPath)
	if err != nil {
		return apiException.RequestError
	}

	// 使用 handleLoginErrors 函数处理响应码
	return handleLoginErrors(resp.Code)
}

// handleRegErrors 根据响应码处理不同的错误
func handleLoginErrors(code int) error {
	switch code {
	case 404:
		return apiException.UserNotFound
	case 405:
		return apiException.NoThatPasswordOrWrong
	case 200:
		return nil
	default:
		return apiException.ServerError
	}
}
