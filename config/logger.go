package config

// desc: 加载日志句柄

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var defLogger = zap.NewExample()

// Logger 配置解析模块
type Logger struct {
	// Output to file or console.
	Mode string

	Level string

	Name, Alone, Path, UUIDPath, ErrPath string

	MaxSize int

	MaxBackups int

	MaxAge int

	Compress bool

	Development bool
}

func newLogger(cfg *Logger) *zap.Logger {
	// 仅打印Error级别以上的日志
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// 获取日志级别
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return defLogger
	}
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})

	errHook := lumberjack.Logger{
		Filename:   cfg.ErrPath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		LocalTime:  true,
	}

	hook := lumberjack.Logger{
		Filename:   cfg.Path,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		LocalTime:  true,
	}

	// 根据环境设置不同日志格式
	var encoder zapcore.Encoder
	if cfg.Development {
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	} else {
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(&hook), lowPriority),
		zapcore.NewCore(encoder, zapcore.AddSync(&errHook), highPriority),
	)

	return zap.New(core)
}
