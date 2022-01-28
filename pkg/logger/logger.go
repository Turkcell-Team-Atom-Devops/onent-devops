package logger

import (
	"go.uber.org/zap"
	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/config"
)

var Log *zap.Logger = New()

func New() *zap.Logger {
	
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	cfg.Encoding = "json"
	cfg.OutputPaths = []string{config.Config.Log.Path}

	var loggerInstance *zap.Logger
	loggerInstance, _ = cfg.Build()
	
	return loggerInstance
}
