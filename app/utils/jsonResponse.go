package utils

import (
	"net/http"
	"runtime"

	"4u-go/app/apiException"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Level 日志级别
type Level uint8

// 日志级别常量
const (
	LevelFatal  Level = 0
	LevelPanic  Level = 1
	LevelDpanic Level = 2
	LevelError  Level = 3
	LevelWarn   Level = 4
	LevelInfo   Level = 5
	LevelDebug  Level = 6
)

// JsonResponse 返回json格式数据
func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data any) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// JsonSuccessResponse 返回成功json格式数据
func JsonSuccessResponse(c *gin.Context, data any) {
	JsonResponse(c, http.StatusOK, 200, "OK", data)
}

// JsonErrorResponse 返回错误json格式数据
func JsonErrorResponse(c *gin.Context, apiErr *apiException.Error, level Level, err error) {
	logError(c, apiErr, level, err)
	JsonResponse(c, http.StatusOK, apiErr.Code, apiErr.Msg, nil)
}

// logError 记录错误日志
func logError(c *gin.Context, apiErr *apiException.Error, level Level, err error) {
	// 获取抛出错误的函数和代码行数
	funcName, file, line := getErrorCallerInfo()
	// 构建日志字段
	logFields := []zap.Field{
		zap.Int("error_code", apiErr.Code),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("func", funcName), // 记录抛出错误的函数名
		zap.String("file", file),     // 记录文件名
		zap.Int("line", line),        // 记录代码行号
		zap.Error(err),               // 记录原始错误信息
	}
	// 记录日志
	switch level {
	case LevelFatal:
		zap.L().Fatal(apiErr.Msg, logFields...)
	case LevelPanic:
		zap.L().Panic(apiErr.Msg, logFields...)
	case LevelDpanic:
		zap.L().DPanic(apiErr.Msg, logFields...)
	case LevelError:
		zap.L().Error(apiErr.Msg, logFields...)
	case LevelWarn:
		zap.L().Warn(apiErr.Msg, logFields...)
	case LevelInfo:
		zap.L().Info(apiErr.Msg, logFields...)
	case LevelDebug:
		zap.L().Debug(apiErr.Msg, logFields...)
	}
}

// getErrorCallerInfo 获取抛出错误的函数名、文件名和行号
func getErrorCallerInfo() (funcName, file string, line int) {
	// 获取调用栈信息，skip 3 层：runtime.Callers, getErrorCallerInfo, logError
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return "unknown", "unknown", 0
	}

	fn := runtime.FuncForPC(pc)
	funcName = fn.Name() // 获取函数名
	return funcName, file, line
}
