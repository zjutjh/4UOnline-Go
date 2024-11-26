package midwares

import (
	"4u-go/app/apiException"
	"4u-go/app/config"
	"github.com/gin-gonic/gin"
)

// CheckInit 中间件用于检查系统是否已初始化。
func CheckInit(c *gin.Context) {
	inited := config.GetInit()
	if !inited {
		apiException.AbortWithException(c, apiException.NotInit, nil)
		return
	}
	c.Next()
}
