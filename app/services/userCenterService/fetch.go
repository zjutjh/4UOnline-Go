package userCenterService

import (
	"encoding/json"

	"4u-go/app/apiException"
	"4u-go/app/utils/fetch"
	"4u-go/config/api/userCenterApi"
)

// UserCenterResponse 用户中心响应结构体
type UserCenterResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// FetchHandleOfPost 向用户中心发送post请求
func FetchHandleOfPost(form map[string]string, url userCenterApi.UserCenterApi) (*UserCenterResponse, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostJsonForm(userCenterApi.UserCenterHost+string(url), form)
	if err != nil {
		return nil, apiException.RequestError
	}
	rc := UserCenterResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiException.RequestError
	}
	return &rc, nil
}
