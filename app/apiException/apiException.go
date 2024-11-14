package apiException

import (
	"net/http"

	"4u-go/app/utils/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Error 表示自定义错误，包括状态码、消息和日志级别。
type Error struct {
	Code  int
	Msg   string
	Level log.Level
}

// Error 表示自定义的错误类型
var (
	ServerError           = NewError(200500, log.LevelError, "系统异常，请稍后重试!")
	OpenIDError           = NewError(200500, log.LevelInfo, "系统异常，请稍后重试!")
	ParamError            = NewError(200501, log.LevelInfo, "参数错误")
	ReactiveError         = NewError(200502, log.LevelInfo, "该通行证已经存在，请重新输入")
	UserAlreadyExisted    = NewError(200503, log.LevelInfo, "该用户已激活")
	RequestError          = NewError(200504, log.LevelInfo, "系统异常，请稍后重试!")
	StudentNumAndIidError = NewError(200505, log.LevelInfo, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError              = NewError(200506, log.LevelInfo, "密码长度必须在6~20位之间")
	UserNotFound          = NewError(200507, log.LevelInfo, "该用户不存在")
	NoThatPasswordOrWrong = NewError(200508, log.LevelInfo, "密码错误")
	NotLogin              = NewError(200509, log.LevelInfo, "未登录")
	NotPermission         = NewError(200510, log.LevelInfo, "该用户无权限")
	ActivityNotFound      = NewError(200511, log.LevelInfo, "活动不存在")
	AnnouncementNotFound  = NewError(200512, log.LevelInfo, "公告不存在")
	AdminKeyError         = NewError(200513, log.LevelInfo, "管理员注册密钥错误")
	AdminAlreadyExisted   = NewError(200514, log.LevelInfo, "管理员账号已存在")
	CollageNotFound       = NewError(200515, log.LevelInfo, "学院不存在")

	NotInit  = NewError(200404, log.LevelWarn, http.StatusText(http.StatusNotFound))
	NotFound = NewError(200404, log.LevelWarn, http.StatusText(http.StatusNotFound))
	Unknown  = NewError(300500, log.LevelError, "系统异常，请稍后重试!")
)

// Error 方法实现了 error 接口，返回错误的消息内容
func (e *Error) Error() string {
	return e.Msg
}

// NewError 创建并返回一个新的自定义错误实例
func NewError(code int, level log.Level, msg string) *Error {
	return &Error{
		Code:  code,
		Msg:   msg,
		Level: level,
	}
}

// AbortWithException 用于返回自定义错误信息
func AbortWithException(c *gin.Context, apiError *Error, err error) {
	logError(c, apiError, err)
	_ = c.AbortWithError(200, apiError) //nolint:errcheck
}

// logError 记录错误日志
func logError(c *gin.Context, apiErr *Error, err error) {
	// 构建日志字段
	logFields := []zap.Field{
		zap.Int("error_code", apiErr.Code),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.Error(err), // 记录原始错误信息
	}
	log.GetLogFunc(apiErr.Level)(apiErr.Msg, logFields...)
}
