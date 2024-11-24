package session

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Init 初始化会话管理，设置会话存储驱动
func Init(r *gin.Engine) error {
	config := getConfig()
	switch config.Driver {
	case Redis:
		return setRedis(r, config.Name)
	case Memory:
		setMemory(r, config.Name)
	default:
		return errors.New("session config error")
	}

	return nil
}
