package utils

import (
	"net/http"

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
	// 构建日志字段
	logFields := []zap.Field{
		zap.Int("error_code", apiErr.Code),
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
		zap.String("ip", c.ClientIP()),
		zap.Error(err), // 记录原始错误信息
	}
	// 创建日志级别映射表
	logMap := map[Level]func(string, ...zap.Field){
		LevelFatal:  zap.L().Fatal,
		LevelPanic:  zap.L().Panic,
		LevelDpanic: zap.L().DPanic,
		LevelError:  zap.L().Error,
		LevelWarn:   zap.L().Warn,
		LevelInfo:   zap.L().Info,
		LevelDebug:  zap.L().Debug,
	}

	// 根据日志级别记录日志
	if logFunc, ok := logMap[level]; ok {
		logFunc(apiErr.Msg, logFields...)
	}
}
