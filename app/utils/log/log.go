package log

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"4u-go/config/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 是应用程序的全局日志记录器
var Logger *zap.Logger

// zapStacktraceMutex 用于保护堆栈跟踪的并发访问
var zapStacktraceMutex sync.Mutex

// Dir 存储日志文件目录
var Dir string

// Config 用于定义日志配置的结构体
type Config struct {
	Development       bool   // 是否开启开发模式
	DisableCaller     bool   // 是否禁用调用方信息
	DisableStacktrace bool   // 是否禁用堆栈跟踪
	Encoding          string // 日志编码格式
	Level             string // 日志级别
	Name              string // 日志名称
	Writers           string // 日志输出方式
	LoggerDir         string // 日志文件目录
	LogRollingPolicy  string // 日志滚动策略
	LogBackupCount    uint   // 日志备份数量
}

// loggerLevelMap 映射日志级别字符串到 zapcore.Level
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

const (
	// WriterConsole 表示控制台输出
	WriterConsole = "console"
	// WriterFile 表示文件输出
	WriterFile = "file"
	// LogSuffix 普通日志后缀
	LogSuffix = ".log"
	// WarnLogSuffix 警告日志后缀
	WarnLogSuffix = "_warn.log"
	// ErrorLogSuffix 错误日志后缀
	ErrorLogSuffix = "_error.log"
)

const (
	// RotateTimeDaily 每日滚动
	RotateTimeDaily = "daily"
	// RotateTimeHourly 每小时滚动
	RotateTimeHourly = "hourly"
)

// loadConfig 加载日志配置
func loadConfig() *Config {
	return &Config{
		Development:       config.Config.GetBool("log.development"),        // 是否是开发环境
		DisableCaller:     config.Config.GetBool("log.disableCaller"),      // 是否禁用调用方
		DisableStacktrace: config.Config.GetBool("log.disableStacktrace"),  // 是否禁用堆栈跟踪
		Encoding:          config.Config.GetString("log.encoding"),         // 编码格式
		Level:             config.Config.GetString("log.level"),            // 日志级别
		Name:              config.Config.GetString("log.name"),             // 日志名称
		Writers:           config.Config.GetString("log.writers"),          // 日志输出方式
		LoggerDir:         config.Config.GetString("log.loggerDir"),        // 日志目录
		LogRollingPolicy:  config.Config.GetString("log.logRollingPolicy"), // 日志滚动策略
		LogBackupCount:    config.Config.GetUint("log.logBackupCount"),     // 日志备份数量
	}
}

// ZapInit 初始化 zap 日志记录器
func ZapInit() {
	cfg := loadConfig()

	Dir = cfg.LoggerDir
	if strings.HasSuffix(Dir, "/") {
		Dir = strings.TrimRight(Dir, "/")
	}

	// 创建日志目录
	if err := createLogDirectory(cfg.LoggerDir); err != nil {
		return
	}

	encoder := createEncoder(cfg)

	var cores []zapcore.Core
	options := []zap.Option{zap.Fields(zap.String("serviceName", cfg.Name))}

	// 根据配置选择输出方式
	cores = append(cores, createLogCores(cfg, encoder, options)...)

	// 合并所有核心
	combinedCore := zapcore.NewTee(cores...)

	// 添加其他选项
	addAdditionalOptions(cfg, &options)

	Logger = zap.New(combinedCore, options...) // 创建新的 zap 日志记录器
	Logger.Info("Logger initialized")          // 初始化日志记录器信息
}

// getLoggerLevel 返回日志级别
func getLoggerLevel(cfg *Config) zapcore.Level {
	level, exist := loggerLevelMap[strings.ToLower(cfg.Level)]
	if !exist {
		return zapcore.DebugLevel // 默认返回 Debug 级别
	}
	return level
}

// getAllCore 返回一个记录所有级别日志的核心
func getAllCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	allWriter := getLogWriterWithTime(cfg, GetLogFile(cfg.Name, LogSuffix))
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel // 记录所有级别到 Fatal
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
}

// getInfoCore 返回一个记录信息级别日志的核心
func getInfoCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	infoWrite := getLogWriterWithTime(cfg, GetLogFile(cfg.Name, LogSuffix))
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel // 记录信息及以上级别日志
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(infoWrite), infoLevel)
}

func getLogCore(encoder zapcore.Encoder, cfg *Config, suffix string, level zapcore.Level) (zapcore.Core, zap.Option) {
	logWrite := getLogWriterWithTime(cfg, GetLogFile(cfg.Name, suffix))
	var stacktrace zap.Option
	logLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if !cfg.DisableCaller {
			zapStacktraceMutex.Lock()
			stacktrace = zap.AddStacktrace(level) // 记录堆栈跟踪
			zapStacktraceMutex.Unlock()
		}
		// 根据传入的日志级别决定记录条件
		if level == zapcore.WarnLevel {
			return lvl == zapcore.WarnLevel // 仅记录警告级别日志
		}
		return lvl >= level // 记录错误及以上级别日志
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(logWrite), logLevel), stacktrace
}

// getLogWriterWithTime 返回一个带时间的日志写入器
func getLogWriterWithTime(cfg *Config, filename string) io.Writer {
	logFullPath := filename
	rotationPolicy := cfg.LogRollingPolicy
	backupCount := cfg.LogBackupCount

	var (
		rotateDuration time.Duration
		timeFormat     string
	)
	// 根据滚动策略设置时间格式
	if rotationPolicy == RotateTimeHourly {
		rotateDuration = time.Hour
		timeFormat = ".%Y%m%d%H"
	} else if rotationPolicy == RotateTimeDaily {
		rotateDuration = time.Hour * 24
		timeFormat = ".%Y%m%d"
	}

	// 检查日志文件是否存在
	if _, err := os.Stat(logFullPath); os.IsNotExist(err) {
		// 如果日志文件不存在，创建它
		if err := createLogFile(logFullPath); err != nil {
			zap.S().Error("Failed to create log file:", err)
			panic(err)
		}
	}

	// 创建轮转日志写入器
	hook, err := rotatelogs.New(
		logFullPath+time.Now().Format(timeFormat),
		rotatelogs.WithLinkName(logFullPath),
		rotatelogs.WithRotationCount(backupCount),
		rotatelogs.WithRotationTime(rotateDuration),
	)

	if err != nil {
		zap.S().Error("Failed to initialize log rotation:", err)
		panic(err)
	}
	return hook
}

// GetLogFile 生成日志文件的完整路径
func GetLogFile(filename string, suffix string) string {
	return ConcatString(config.Config.GetString("log.loggerDir"), "/", filename, suffix)
}

// ConcatString 将多个字符串连接成一个字符串
func ConcatString(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		// 检查 WriteString 返回的错误，尽管不太可能返回错误
		if _, err := buffer.WriteString(i); err != nil {
			// 这里可以选择记录错误或者其他处理，取决于您的需求
			zap.S().Error("Failed to write string to buffer:", err)
		}
	}
	return buffer.String()
}

// createLogFile 创建日志文件并设置权限
func createLogFile(logFullPath string) error {
	file, err := os.Create(logFullPath) //nolint:gosec
	if err != nil {
		return err // 返回错误，而不是 panic
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			zap.S().Error("Failed to close log file:", closeErr)
		}
	}()

	// 设置日志文件权限为 0600
	return os.Chmod(logFullPath, 0600)
}

// createLogDirectory 创建日志目录
func createLogDirectory(dir string) error {
	if err := os.MkdirAll(dir, 0750); err != nil {
		zap.S().Error("创建日志目录失败:", err)
		return err
	}
	return nil
}

// createEncoder 创建日志编码器
func createEncoder(cfg *Config) zapcore.Encoder {
	var encoderCfg zapcore.EncoderConfig
	if cfg.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	// 自定义字段名称
	encoderCfg.LevelKey = "level"                      // 原来的 "L"
	encoderCfg.TimeKey = "timestamp"                   // 原来的 "T"
	encoderCfg.CallerKey = "caller"                    // 原来的 "C"
	encoderCfg.MessageKey = "message"                  // 原来的 "M"
	encoderCfg.StacktraceKey = "stacktrace"            // 原来的堆栈跟踪
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder // 设置时间编码格式

	if cfg.Encoding == WriterConsole {
		return zapcore.NewConsoleEncoder(encoderCfg) // 控制台编码器
	}
	return zapcore.NewJSONEncoder(encoderCfg) // JSON 编码器
}

// addAdditionalOptions 添加额外的选项
func addAdditionalOptions(cfg *Config, options *[]zap.Option) {
	if !cfg.DisableCaller {
		*options = append(*options, zap.AddCaller()) // 添加调用方信息
	}
	if !cfg.DisableStacktrace {
		*options = append(*options, zap.AddStacktrace(zapcore.ErrorLevel)) // 添加堆栈跟踪
	}
}

// createLogCores 创建日志核心
func createLogCores(cfg *Config, encoder zapcore.Encoder, options []zap.Option) []zapcore.Core {
	var cores []zapcore.Core
	writers := strings.Split(cfg.Writers, ",")

	for _, writer := range writers {
		cores = append(cores, createCoreForWriter(writer, encoder, cfg, options)...)
	}
	return cores
}

// createCoreForWriter 根据不同的 writer 类型创建相应的日志核心
func createCoreForWriter(writer string, encoder zapcore.Encoder, cfg *Config, options []zap.Option) []zapcore.Core {
	switch writer {
	case WriterConsole:
		return []zapcore.Core{
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)),
		}
	case WriterFile:
		return createFileCores(encoder, cfg, options)
	default:
		return []zapcore.Core{
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)),
			getAllCore(encoder, cfg),
		}
	}
}

// createFileCores 创建文件相关的日志核心
func createFileCores(encoder zapcore.Encoder, cfg *Config, options []zap.Option) []zapcore.Core {
	var cores []zapcore.Core
	cores = append(cores, getInfoCore(encoder, cfg))

	// 获取警告日志核心
	if warnCore, warnOption := getLogCore(encoder, cfg, WarnLogSuffix, zapcore.WarnLevel); warnCore != nil {
		cores = append(cores, warnCore)
		if warnOption != nil {
			options = append(options, warnOption) //nolint:revive
		}
	}

	// 获取错误日志核心
	if errorCore, errorOption := getLogCore(encoder, cfg, ErrorLogSuffix, zapcore.ErrorLevel); errorCore != nil {
		cores = append(cores, errorCore)
		if errorOption != nil {
			options = append(options, errorOption) //nolint:revive
		}
	}
	return cores
}
