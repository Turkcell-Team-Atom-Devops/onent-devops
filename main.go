package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/config"
	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/logger"
	"github.com/Turkcell-Team-Atom-Devops/onent-devops/pkg/router"
	"go.uber.org/zap"
)

func main() {
	srv := &http.Server{
		Addr:         ":" + config.Config.Server.Port,
		WriteTimeout: time.Second * time.Duration(config.Config.Server.Timeout),
		ReadTimeout:  time.Second * time.Duration(config.Config.Server.Timeout),
		IdleTimeout:  time.Second * 60,
		Handler:      router.New(),
	}

	logger.Log.Info("listening on ", zap.String("port", config.Config.Server.Port))
	if err := srv.ListenAndServe(); err != nil {
		logger.Log.Error(fmt.Sprintf("failed to start server: %v", err))
	}
}
