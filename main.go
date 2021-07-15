package main

import (
	"time"
	"zap_di/component"
	"zap_di/env"
	"zap_di/logger"

	"go.uber.org/zap"
)

func main() {
	env.SetEnv()
	err := logger.SetUp()
	if err != nil {
		panic(err)
	}
	logger.Info("Zap Start.", zap.Time("time", time.Now()))
	logger.InfoQuick("Zap Start.", zap.Time("time", time.Now()))
	logger.Debug("Debug.", zap.Time("time", time.Now()))
	component.LogTest()
}
