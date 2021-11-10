package component

import (
	"time"
	"zap_di/logger"

	"go.uber.org/zap"
)

func LogTest() {
	logger.Info("Test Check time.", zap.Time("time", time.Now()))
}
