package midwares

import (
	"errors"
	"fmt"
	"log"

	"4u-go/app/apiException"
	"github.com/gin-gonic/gin"
)

// ErrHandler 中间件用于处理请求错误。
// 如果存在错误，将返回相应的 JSON 响应。
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()        // 继续处理请求
		handleErrors(c) // 处理可能的错误
	}
}

// handleErrors 处理上下文中的错误并返回相应的 JSON 响应。
func handleErrors(c *gin.Context) {
	if length := len(c.Errors); length > 0 {
		e := c.Errors[length-1] // 获取最后一个错误
		err := e.Err
		var apiErr *apiException.Error

		// 根据错误类型进行处理
		apiErr = getAPIError(err)

		if apiErr != nil {
			c.JSON(apiErr.StatusCode, apiErr) // 返回相应的错误响应
			return
		}

		// 打印错误日志
		if _, err := fmt.Println(c.Errors); err != nil {
			log.Println("Error printing errors:", err) // 处理 fmt.Println 的潜在错误
		}
	}
}

// getAPIError 根据不同的错误类型返回相应的 apiException.Error。
func getAPIError(err error) *apiException.Error {
	if err == nil {
		return nil
	}

	var apiErr *apiException.Error
	if errors.As(err, &apiErr) {
		return apiErr
	}
	return apiException.OtherError(err.Error()) // 如果不是自定义错误，则返回其他错误
}

// HandleNotFound 处理 404 错误。
func HandleNotFound(c *gin.Context) {
	err := apiException.NotFound
	c.JSON(err.StatusCode, err)
}
