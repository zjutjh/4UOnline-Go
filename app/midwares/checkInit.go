package midwares

import (
	"4u-go/app/apiException"
	"4u-go/app/config"
	"github.com/gin-gonic/gin"
)

func CheckInit(c *gin.Context) {
	inited := config.GetInit()
	if !inited {
		_ = c.AbortWithError(200, apiException.NotInit)
		return
	}
	c.Next()
}
