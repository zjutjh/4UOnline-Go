package apiException

import "net/http"

// Error 表示自定义错误，包括状态码、代码和消息。
type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

// Error 表示自定义的错误类型
var (
	ServerError = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	OpenIDError = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	ParamError  = NewError(http.StatusInternalServerError, 200501, "参数错误")

	NotInit  = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	NotFound = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	Unknown  = NewError(http.StatusInternalServerError, 300500, "系统异常，请稍后重试!")
)

// OtherError 返回一个表示其他未定义错误的自定义错误消息
func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

// Error 方法实现了 error 接口，返回错误的消息内容
func (e *Error) Error() string {
	return e.Msg
}

// NewError 创建并返回一个新的自定义错误实例
func NewError(statusCode, code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
	}
}
