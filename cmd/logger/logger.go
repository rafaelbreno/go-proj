package logger

import (
	"go.uber.org/zap"
)

func Log() *zap.Logger {
	Log, err := zap.NewProduction()

	if err != nil {
		panic(err)
	}

	Log.Info("Starting application")

	return Log
}
