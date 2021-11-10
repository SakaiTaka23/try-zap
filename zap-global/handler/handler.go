package handler

import (
	"go.uber.org/zap"
)

type resStruct struct {
	response string
	id       int
}

func Hello() {
	testStruct := resStruct{
		response: "Hello World",
		id:       1,
	}
	response := "Hello World"
	id := 1
	zap.L().Info("sending string", zap.String("response", response), zap.Int("number", id))
	zap.S().Infow("sending string ", "response", response, "number", id, "struct", testStruct)
	zap.S().Info("msg", "sending string", "response", response, "number", id)

	zap.S().Debugw("debug struct", "struct", testStruct)
	zap.S().Debug("struct", testStruct)
	zap.S().Debug("msg: ", "debug struct", "struct: ", testStruct)

	zap.S().Errorw("some error occurred")
}
