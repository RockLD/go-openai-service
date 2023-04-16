package conf

import (
	"go.uber.org/zap"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) InitLogger(env string) *zap.Logger {
	var logger *zap.Logger
	var err error
	switch env {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		panic("env error")
	}
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	return logger
}
