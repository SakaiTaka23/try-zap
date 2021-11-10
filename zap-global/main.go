package main

import (
	"zap-global/handler"

	"go.uber.org/zap"
)

func main() {
	production := true
	outputPath := "./logs/test.log"
	var logger *zap.Logger
	if production {
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{outputPath}
		logger, _ = config.Build()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	handler.Hello()
}
