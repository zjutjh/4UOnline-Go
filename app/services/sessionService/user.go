package sessionService

import (
	"errors"

	"4u-go/app/models"
	"4u-go/app/services/userService"
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

// UpdateUserSession 更新用户session缓存
func UpdateUserSession(c *gin.Context) (*models.User, error) {
	user, err := GetUserSession(c)
	if err != nil {
		return nil, err
	}
	err = SetUserSession(c, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserSession 检查用户session缓存
func GetUserSession(c *gin.Context) (*models.User, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return nil, errors.New("")
	}
	user, err := userService.GetUserByID(id.(int))
	if err != nil {
		if err := ClearUserSession(c); err != nil {
			return nil, err
		}
		return nil, errors.New("")
	}
	return user, nil
}

// ClearUserSession 清空用户session缓存
func ClearUserSession(c *gin.Context) error {
	webSession := sessions.Default(c)
	webSession.Delete("id")
	err := webSession.Save()
	return err
}
