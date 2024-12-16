package userCenterService

import (
	"4u-go/app/apiException"
	"4u-go/app/utils/request"
	"4u-go/config/api/userCenterApi"
)

// UserCenterResponse 用户中心响应结构体
type UserCenterResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// FetchHandleOfPost 向用户中心发送 POST 请求
func FetchHandleOfPost(form map[string]any, url string) (*UserCenterResponse, error) {
	client := request.NewUnSafe()
	var rc UserCenterResponse

	// 发送 POST 请求并自动解析 JSON 响应
	resp, err := client.Request().
		SetHeader("Content-Type", "application/json").
		SetBody(form).
		SetResult(&rc).
		Post(userCenterApi.UserCenterHost + url)

	// 检查请求错误
	if err != nil || resp.IsError() {
		return nil, apiException.RequestError
	}

	// 返回解析后的响应
	return &rc, nil
}
