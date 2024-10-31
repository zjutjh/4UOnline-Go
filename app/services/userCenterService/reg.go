package userCenterService

import (
	"net/url"

	"4u-go/app/apiException"
	"4u-go/config/api/userCenterApi"
)

// RegWithoutVerify 用户中心不验证激活用户
func RegWithoutVerify(stuId string, pass string, iid string, email string, userType uint) error {
	params := url.Values{}
	userUrl, err := url.Parse(string(userCenterApi.UCRegWithoutVerify))
	if err != nil {
		return err
	}
	userUrl.RawQuery = params.Encode()
	urlPath := userUrl.String()
	regMap := make(map[string]any)
	regMap["stu_id"] = stuId
	regMap["password"] = pass
	regMap["iid"] = iid
	regMap["email"] = email
	regMap["type"] = userType
	regMap["bound_system"] = 1
	resp, err := FetchHandleOfPost(regMap, userCenterApi.UserCenterApi(urlPath))
	if err != nil {
		return err
	}
	return handleRegErrors(resp.Code)
}

// handleRegErrors 根据响应码处理不同的错误
func handleRegErrors(code int) error {
	switch code {
	case 400, 402:
		return apiException.StudentNumAndIidError
	case 401:
		return apiException.PwdError
	case 403:
		return apiException.ReactiveError
	default:
		return nil
	}
}
