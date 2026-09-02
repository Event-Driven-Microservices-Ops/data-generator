package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/urbaniakmichal/data-generator/internal/api"
	"github.com/urbaniakmichal/data-generator/internal/config"
)

func main() {
	cfg := loadConfig("internal/config/config.yaml")

	server := api.NewServer(cfg)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting HTTP server...", slog.String("port", cfg.Port))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("HTTP server failed", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func loadConfig(path string) *config.ServerConfig {
	cfg, err := config.Load(path)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return cfg
}
