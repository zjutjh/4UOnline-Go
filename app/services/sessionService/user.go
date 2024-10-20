package sessionService

import (
	"4u-go/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetUserSession 设置用户session缓存
func SetUserSession(c *gin.Context, user *models.User) error {
	webSession := sessions.Default(c)
	webSession.Options(sessions.Options{MaxAge: 3600 * 24 * 7, Path: "/api"})
	webSession.Set("id", user.ID)
	return webSession.Save()
}
