package midwares

import (
	"log"

	"4u-go/app/apiException"
	"4u-go/app/config"
	"github.com/gin-gonic/gin"
)

// CheckInit 中间件用于检查系统是否已初始化。
func CheckInit(c *gin.Context) {
	inited := config.GetInit()
	if !inited {
		err := c.AbortWithError(200, apiException.NotInit)
		if err != nil {
			log.Println("CheckInitFailed:", err) // 记录错误日志，而不是退出程序
		}
		return
	}
	c.Next()
}
