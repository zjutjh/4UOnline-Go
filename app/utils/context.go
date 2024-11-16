package utils

import (
	"4u-go/app/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUser 从上下文中提取 *models.User
func GetUser(c *gin.Context) *models.User {
	if val, ok := c.Get("user"); ok {
		if user, ok := val.(*models.User); ok {
			return user
		}
	}
	zap.L().Error("从上下文中提取 *models.User 失败")
	return &models.User{}
}
