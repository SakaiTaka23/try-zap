package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	sink := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/zap.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     10,
		Compress:   true,
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		sink,
		zap.InfoLevel,
	)
	logger := zap.New(core)
	logger.Info("hoge")
}
