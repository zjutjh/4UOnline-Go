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
	ServerError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	OpenIDError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	ParamError            = NewError(http.StatusInternalServerError, 200501, "参数错误")
	ReactiveError         = NewError(http.StatusInternalServerError, 200502, "该通行证已经存在，请重新输入")
	UserAlreadyExisted    = NewError(http.StatusInternalServerError, 200503, "该用户已激活")
	RequestError          = NewError(http.StatusInternalServerError, 200504, "系统异常，请稍后重试!")
	StudentNumAndIidError = NewError(http.StatusInternalServerError, 200505, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError              = NewError(http.StatusInternalServerError, 200506, "密码长度必须在6~20位之间")
	UserNotFound          = NewError(http.StatusInternalServerError, 200507, "该用户不存在")
	NoThatPasswordOrWrong = NewError(http.StatusInternalServerError, 200508, "密码错误")
	NotLogin              = NewError(http.StatusInternalServerError, 200509, "未登录")
	NotPermission         = NewError(http.StatusInternalServerError, 200510, "该用户无权限")
	ActivityNotFound      = NewError(http.StatusInternalServerError, 200511, "活动不存在")

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
