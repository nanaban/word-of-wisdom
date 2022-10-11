package common

import (
	"log"

	"go.uber.org/zap"
)

func MustNewLogger(isDebug bool) (logger *zap.Logger) {
	var err error
	if isDebug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	return logger
}
