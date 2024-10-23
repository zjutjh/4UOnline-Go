package utils

import (
	"net/http"

	"4u-go/app/apiException"
	"github.com/gin-gonic/gin"
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
func JsonErrorResponse(c *gin.Context, err *apiException.Error) {
	JsonResponse(c, http.StatusOK, err.Code, err.Msg, nil)
}