package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	Log    *zap.Logger
	Config zap.Config
}

var log *zap.Logger

func Innit() {
	var l Logger

	l.setConfig()

	l.log()

	log = l.Log

	defer l.Log.Info("Starting application...")
}

func (l *Logger) setConfig() {
	l.Config = zap.NewProductionConfig()

	encConfig := zap.NewProductionEncoderConfig()
	encConfig.TimeKey = "timestamp"
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encConfig.StacktraceKey = ""
	l.Config.EncoderConfig = encConfig
}

func (l *Logger) log() {

	log, err := l.Config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}

	l.Log = log
}

func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}
