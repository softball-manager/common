package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	InfoLevel  = "info"
	DebugLevel = "debug"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"

	logLevels = map[string]zapcore.Level{
		DebugLevel: zap.DebugLevel,
		InfoLevel:  zap.InfoLevel,
		WarnLevel:  zap.WarnLevel,
		ErrorLevel: zap.ErrorLevel,
		PanicLevel: zap.PanicLevel,
	}

	TableNameLogKey = "table name"
	PlayerIDLogKey  = "pid"
	TeamIDLogKey    = "tid"
	GameIDLogKey    = "gid"
	EnvLogKey       = "env"
)

func GetLogger(level string) *zap.Logger {
	logger := newLogger(level)
	defer logger.Sync()
	return logger
}

func GetLoggerWithEnv(level string, env string) *zap.Logger {
	logger := newLogger(level).With(zap.String(EnvLogKey, env))
	defer logger.Sync()
	return logger
}

func newLogger(level string) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevels[level]),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	return zap.Must(config.Build())
}
