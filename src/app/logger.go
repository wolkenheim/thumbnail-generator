package app

import (
	"encoding/json"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
    zapConfig := getCustomLoggerConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals (logger)

	// call sugared or regular logger globally with
	// zap.S().Info("message goes here")
}

func NewZapLogger() *zap.Logger{
	zapConfig := getCustomLoggerConfig()
	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}

func getCustomLoggerConfig() zap.Config{
	jsonInit := []byte(`{"classtype": "application"}`)
	var initMap map[string]interface{}
	json.Unmarshal(jsonInit, &initMap)

	zapConfig := zap.NewProductionConfig()
	zapConfig.Encoding = "json" // json | console
	zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	zapConfig.DisableStacktrace = true
	zapConfig.InitialFields = initMap

	return zapConfig
}