package midwares

import (
	"4u-go/app/apiException"
	"4u-go/app/services/sessionService"
	"github.com/gin-gonic/gin"
)

// CheckAdmin 验证管理员权限
func CheckAdmin(c *gin.Context) {
	user, err := sessionService.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	if user.Type < 2 { // 验证管理员权限
		apiException.AbortWithException(c, apiException.NotPermission, nil)
		return
	}
	c.Set("admin_type", user.Type)
	c.Set("user_id", user.ID)
	c.Next()
}
