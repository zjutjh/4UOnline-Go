package userCenterService

import (
	"net/url"

	"4u-go/app/apiException"
	"4u-go/config/api/userCenterApi"
)

// DeleteAccount 注销账户
func DeleteAccount(stuid, iid string) error {
	deleteUrl, err := url.Parse(userCenterApi.DelAccount)
	if err != nil {
		return err
	}
	urlPath := deleteUrl.String()
	regMap := map[string]any{
		"iid":          iid,
		"stuid":        stuid,
		"bound_system": 1,
	}
	resp, err := FetchHandleOfPost(regMap, urlPath)
	if err != nil {
		return err
	}
	return handleDeleteErrors(resp.Code)
}

// handleDeleteErrors 根据响应码处理不同的错误
func handleDeleteErrors(code int) error {
	switch code {
	case 400:
		return apiException.StudentNumAndIidError
	case 404:
		return apiException.UserNotFound
	case 200:
		return nil
	default:
		return apiException.ServerError
	}
}
