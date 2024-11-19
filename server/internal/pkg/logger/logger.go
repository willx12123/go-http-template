package logger

import (
	"go.uber.org/zap"

	"server/internal/pkg/config"
)

var Default *zap.Logger

func Init() {
	var (
		logger *zap.Logger
		err    error
	)
	if config.IsProd() {
		logger, err = zap.NewProduction(zap.WithCaller(false))
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	Default = logger
}
