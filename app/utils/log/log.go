package log

import (
	"io"
	"os"
	"strings"

	"4u-go/config/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

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
	LogMaxSize        int    // 日志文件最大大小（单位：MB）
	LogMaxAge         int    // 日志文件最大保存天数
	LogCompress       bool   // 是否压缩日志
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
)

// loadConfig 加载日志配置
func loadConfig() *Config {
	return &Config{
		Development:       config.Config.GetBool("log.development"),       // 是否是开发环境
		DisableCaller:     config.Config.GetBool("log.disableCaller"),     // 是否禁用调用方
		DisableStacktrace: config.Config.GetBool("log.disableStacktrace"), // 是否禁用堆栈跟踪
		Encoding:          config.Config.GetString("log.encoding"),        // 编码格式
		Level:             config.Config.GetString("log.level"),           // 日志级别
		Name:              config.Config.GetString("log.name"),            // 日志名称
		Writers:           config.Config.GetString("log.writers"),         // 日志输出方式
		LoggerDir:         config.Config.GetString("log.loggerDir"),       // 日志目录
		LogCompress:       config.Config.GetBool("log.logCompress"),       // 是否压缩日志
		LogMaxSize:        config.Config.GetInt("log.logMaxSize"),         // 日志文件最大大小（单位：MB）
		LogMaxAge:         config.Config.GetInt("log.logMaxAge"),          // 日志保存天数
	}
}

// ZapInit 初始化 zap 日志记录器
func ZapInit() {
	cfg := loadConfig()

	Dir = cfg.LoggerDir
	if strings.HasSuffix(Dir, "/") {
		Dir = strings.TrimRight(Dir, "/") // 去除尾部斜杠
	}

	// 创建日志目录
	if err := createLogDirectory(Dir); err != nil {
		return
	}

	encoder := createEncoder(cfg)

	var cores []zapcore.Core
	options := []zap.Option{zap.Fields(zap.String("serviceName", cfg.Name))}

	// 根据配置选择输出方式
	cores = append(cores, createLogCores(cfg, encoder)...)

	// 合并所有核心
	combinedCore := zapcore.NewTee(cores...)

	// 添加其他选项
	addAdditionalOptions(cfg, &options)

	logger := zap.New(combinedCore, options...) // 创建新的 zap 日志记录器
	zap.ReplaceGlobals(logger)                  // 替换全局日志记录器

	zap.L().Info("Logger initialized") // 初始化日志记录器信息
}

// getLoggerLevel 返回日志级别
func getLoggerLevel(cfg *Config) zapcore.Level {
	level, exist := loggerLevelMap[strings.ToLower(cfg.Level)]
	if !exist {
		return zapcore.DebugLevel // 默认返回 Debug 级别
	}
	return level
}

// getFileCore 返回一个把所有级别日志输出到文件的核心
func getFileCore(encoder zapcore.Encoder, cfg *Config) zapcore.Core {
	allWriter := getLogWriter(cfg, GetLogFilepath(cfg.Name, LogSuffix))
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel // 记录所有级别到 Fatal
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(allWriter), allLevel)
}

// getLogWriter 返回一个日志写入器
func getLogWriter(cfg *Config, filename string) io.Writer {
	return &lumberjack.Logger{
		Filename: filename,
		MaxSize:  cfg.LogMaxSize,  // 最大日志文件大小（单位：MB），可以根据需求配置
		MaxAge:   cfg.LogMaxAge,   // 文件保存的最大天数
		Compress: cfg.LogCompress, // 是否压缩日志
	}
}

// GetLogFilepath 生成日志文件的完整路径
func GetLogFilepath(filename string, suffix string) string {
	return Dir + "/" + filename + suffix
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
	encoderCfg.StacktraceKey = "stacktrace"            // 原来的 "S"
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder // 设置时间编码格式

	if cfg.Encoding == WriterConsole {
		return zapcore.NewConsoleEncoder(encoderCfg) // 控制台编码器
	}
	return zapcore.NewJSONEncoder(encoderCfg) // JSON 编码器
}

// addAdditionalOptions 添加额外的选项
func addAdditionalOptions(cfg *Config, options *[]zap.Option) {
	if !cfg.DisableCaller {
		*options = append(*options, zap.AddCallerSkip(2)) // 添加调用方要求
		*options = append(*options, zap.AddCaller())      // 添加调用方信息
	}
	if !cfg.DisableStacktrace {
		*options = append(*options, zap.AddStacktrace(zapcore.ErrorLevel)) // 添加堆栈跟踪
	}
}

// createLogCores 创建日志核心
func createLogCores(cfg *Config, encoder zapcore.Encoder) []zapcore.Core {
	var cores []zapcore.Core
	writers := strings.Split(cfg.Writers, ",")

	for _, writer := range writers {
		switch writer {
		case WriterConsole:
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(cfg)))
		case WriterFile:
			cores = append(cores, getFileCore(encoder, cfg))
		}
	}
	return cores
}
