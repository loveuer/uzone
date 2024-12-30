package log

import (
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	config = zap.NewProductionConfig()
)

func init() {
	var err error
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

type Level zapcore.Level

const (
	LevelDebug Level = Level(zapcore.DebugLevel)
	LevelInfo  Level = Level(zapcore.InfoLevel)
	LevelWarn  Level = Level(zapcore.WarnLevel)
	LevelError Level = Level(zapcore.ErrorLevel)
)

func SetLogLevel(l Level) {
	config.Level = zap.NewAtomicLevelAt(zapcore.Level(l))
}
