package session

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Init 初始化会话管理，设置会话存储驱动
func Init(r *gin.Engine) error {
	config := getConfig()
	switch config.Driver {
	case string(Redis):
		return setRedis(r, config.Name)
	case string(Memory):
		setMemory(r, config.Name)
	default:
		return fmt.Errorf("session configError")
	}

	return nil
}
