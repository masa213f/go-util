package util

import (
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogEncodingJSON    = "json"
	LogEncodingConsole = "console"
)

const (
	LogLevelInfo  = 1
	LogLevelDebug = 1
)

func NewLogger(encoding string, debug bool) logr.Logger {
	lv := zap.InfoLevel
	if debug {
		lv = zap.DebugLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		NameKey:    "logger",
		LevelKey:   "lv",
		TimeKey:    "ts",
		MessageKey: "msg",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02T15:04:05.000000Z0700")) // RFC3339Micro
		},
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}
	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(lv),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Encoding:          encoding,
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}

	zapLog, err := logConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("failed create a logger (%v)", err))
	}
	return zapr.NewLogger(zapLog)
}
