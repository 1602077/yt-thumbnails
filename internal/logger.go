package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	ERROR = "ERROR"
	DEBUG = "DEBUG"
	INFO  = "INFO"
)

var logger *zap.Logger

func Init(level string, encoding string) {
	loggerConfig := zap.NewProductionConfig()

	loggerConfig.DisableStacktrace = false
	loggerConfig.Encoding = encoding

	switch level {
	case INFO:
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case DEBUG:
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	default:
		log.Print("INFO|logger defaulting to a level of 'ERROR'")
		loggerConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}

	// In cloud environments this should be JSON.
	loggerConfig.EncoderConfig.ConsoleSeparator = "|"
	loggerConfig.EncoderConfig.CallerKey = ""
	loggerConfig.EncoderConfig.FunctionKey = ""
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	logger, err = loggerConfig.Build()
	if err != nil {
		panic("failed to initialise logger")
	}
	zap.ReplaceGlobals(logger)
}

func Error(format string, a ...interface{}) {
	zap.S().Errorw(format, a...)
}

func Warn(format string, a ...any) {
	zap.S().Warnw(format, a...)
}

func Debug(format string, a ...any) {
	zap.S().Debugw(format, a...)
}

func Info(format string, a ...any) {
	zap.S().Infow(format, a...)
}

func Fatal(format string, a ...any) {
	zap.S().Fatalw(format, a...)
}
