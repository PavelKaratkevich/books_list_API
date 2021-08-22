package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	//"time"
)


var logger *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	logger, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

func Error(message string, fields ...zap.Field) {
	logger.Error(message, fields...)
}

func Info(message string, fields ...zap.Field) {
	logger.Info(message, fields...)
}
