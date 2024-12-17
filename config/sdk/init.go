package sdk

import (
	"strings"

	"4u-go/config/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zjutjh/WeJH-SDK/aesHelper"
	"github.com/zjutjh/WeJH-SDK/minioHelper"
	"github.com/zjutjh/WeJH-SDK/redisHelper"
	"github.com/zjutjh/WeJH-SDK/sessionHelper"
	"github.com/zjutjh/WeJH-SDK/wechatHelper"
	"github.com/zjutjh/WeJH-SDK/zapHelper"
	"go.uber.org/zap"
)

// RedisClient 全局 Redis 客户端实例
var RedisClient *redis.Client

// MiniProgram 是一个指向小程序实例的指针
var MiniProgram *miniprogram.MiniProgram

// MinioService 全局 Minio 服务实例
var MinioService *minioHelper.Service

// ZapInit 初始化 Zap 日志库
func ZapInit() error {
	zapInfo := zapHelper.InfoConfig{
		StacktraceLevel:   "warn",
		DisableStacktrace: config.Config.GetBool("log.disableStacktrace"), // 是否禁用堆栈跟踪
		ConsoleLevel:      config.Config.GetString("log.level"),           // 日志级别
		Name:              config.Config.GetString("log.name"),            // 日志名称
		Writer:            config.Config.GetString("log.writer"),          // 日志输出方式
		LoggerDir:         config.Config.GetString("log.loggerDir"),       // 日志目录
		LogCompress:       config.Config.GetBool("log.logCompress"),       // 是否压缩日志
		LogMaxSize:        config.Config.GetInt("log.logMaxSize"),         // 日志文件最大大小（单位：MB）
		LogMaxAge:         config.Config.GetInt("log.logMaxAge"),          // 日志保存天数
	}
	logger, err := zapHelper.Init(&zapInfo)
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(logger)
	zap.L().Info("Logger initialized")
	return nil
}

// Init 初始化配置
func Init(r *gin.Engine) error {
	// 初始化 AES
	err := aesHelper.Init(config.Config.GetString("aes.encryptKey"))
	if err != nil {
		return err
	}

	// 初始化 MinIO
	minioInfo := minioHelper.InfoConfig{
		EndPoint:  config.Config.GetString("minio.endPoint"),
		AccessKey: config.Config.GetString("minio.accessKey"),
		SecretKey: config.Config.GetString("minio.secretKey"),
		Secure:    config.Config.GetBool("minio.secure"),
		Bucket:    config.Config.GetString("minio.bucket"),
		Domain:    config.Config.GetString("minio.domain"),
		TempDir:   strings.Trim(config.Config.GetString("minio.tempDir"), " /") + "/",
	}
	MinioService, err = minioHelper.Init(&minioInfo)
	if err != nil {
		return err
	}

	// 初始化 Redis
	redisInfo := redisHelper.InfoConfig{
		Host:     config.Config.GetString("redis.host"),
		Port:     config.Config.GetString("redis.port"),
		DB:       config.Config.GetInt("redis.db"),
		Password: config.Config.GetString("redis.pass"),
	}
	RedisClient = redisHelper.Init(&redisInfo)

	// 初始化会话管理
	sessionInfo := sessionHelper.InfoConfig{
		Name:        config.Config.GetString("session.name"),
		SecretKey:   config.Config.GetString("session.secret"),
		RedisConfig: &redisInfo,
	}
	err = sessionHelper.Init(&sessionInfo, r)
	if err != nil {
		return err
	}

	// 初始化微信小程序
	wechatInfo := wechatHelper.InfoConfig{
		AppId:       config.Config.GetString("wechat.appid"),
		AppSecret:   config.Config.GetString("wechat.appsecret"),
		RedisConfig: &redisInfo,
	}
	MiniProgram = wechatHelper.Init(&wechatInfo)

	return nil
}
