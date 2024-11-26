package userCenterService

import (
	"net/url"

	"4u-go/app/apiException"
	"4u-go/config/api/userCenterApi"
)

// RePassWithoutEmail 不通过邮箱修改密码
func RePassWithoutEmail(stuid, iid, pwd string) error {
	repassUrl, err := url.Parse(userCenterApi.RePassWithoutEmail)
	if err != nil {
		return err
	}
	urlPath := repassUrl.String()
	regMap := map[string]any{
		"stuid": stuid,
		"iid":   iid,
		"pwd":   pwd,
	}
	resp, err := FetchHandleOfPost(regMap, urlPath)
	if err != nil {
		return err
	}
	return handleRePassErrors(resp.Code)
}

// handleRePassErrors 根据响应码处理不同的错误
func handleRePassErrors(code int) error {
	switch code {
	case 400:
		return apiException.StudentNumAndIidError
	case 401:
		return apiException.PwdError
	case 404:
		return apiException.UserNotFound
	case 200:
		return nil
	default:
		return apiException.ServerError
	}
}
