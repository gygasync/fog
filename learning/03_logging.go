package learning

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Getinfo() string {
	return "A year is on average 365.25 days long."
}

func Getwarning() string {
	return "Tornado worning at " + time.Now().String()
}

func Geterror() string {
	return "Error at directory " + os.Getenv("PWD")
}

func Getlogger(color bool) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	if color {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	logger, _ := config.Build()
	return logger
}
