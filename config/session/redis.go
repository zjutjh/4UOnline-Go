package session

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type redisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func setRedis(r *gin.Engine, name string) error {
	info := getRedisConfig()
	store, err := sessionRedis.NewStore(10, "tcp", info.Host+":"+info.Port, info.Password, []byte("secret"))
	if err != nil {
		return fmt.Errorf("redis session init failed: %w", err) // 返回包装后的错误
	}
	r.Use(sessions.Sessions(name, store))
	return nil
}
