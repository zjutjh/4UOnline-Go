package midwares

import (
	"4u-go/app/apiException"
	"4u-go/app/services/sessionService"
	"4u-go/app/utils"
	"github.com/gin-gonic/gin"
)

// CheckAdmin 验证管理员权限
func CheckAdmin(c *gin.Context) {
	user, err := sessionService.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, apiException.NotLogin, utils.LevelInfo, err)
		c.Abort()
		return
	}
	if user.Type < 2 { // 验证管理员权限
		utils.JsonErrorResponse(c, apiException.NotPermission, utils.LevelInfo, nil)
		c.Abort()
		return
	}
	c.Next()
}
