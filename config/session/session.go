package session

import (
	"fmt"
	"strconv"

	"4u-go/config/config"
	"4u-go/config/redis"
	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Init 使用 Redis 初始化会话管理
func Init(r *gin.Engine) error {
	info := redis.InfoConfig
	name := config.Config.GetString("session.name")
	secret := config.Config.GetString("session.secret")

	store, err := sessionRedis.NewStoreWithDB(10, "tcp",
		info.Host+":"+info.Port, info.Password,
		strconv.Itoa(info.DB),
		[]byte(secret),
	)
	if err != nil {
		return fmt.Errorf("session init failed: %w", err)
	}
	r.Use(sessions.Sessions(name, store))
	return nil
}
