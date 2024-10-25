package midwares

import (
	"net/http"
	"runtime"

	"4u-go/app/apiException"
	"4u-go/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrHandler 中间件用于处理请求错误。
// 如果存在错误，将返回相应的 JSON 响应。
func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 恢复恐慌并处理
			if err := recover(); err != nil {
				// 打印恐慌信息及堆栈跟踪
				stackTrace := make([]byte, 4096)
				stackSize := runtime.Stack(stackTrace, true)
				// After
				zap.L().Panic("Panic recovered",
					zap.Any("error", err),
					zap.ByteString("stackTrace", stackTrace[:stackSize]))

				// 返回 500 错误响应
				apiErr := apiException.ServerError
				utils.JsonResponse(c, http.StatusOK, apiErr.Code, apiErr.Msg, nil)
			}
		}()

		c.Next() // 继续处理请求
	}
}

// HandleNotFound 处理 404 错误。
func HandleNotFound(c *gin.Context) {
	err := apiException.NotFound
	// 记录 404 错误日志
	zap.L().Warn("404 Not Found",
		zap.String("path", c.Request.URL.Path),
		zap.String("method", c.Request.Method),
	)
	c.JSON(err.StatusCode, err)
}
