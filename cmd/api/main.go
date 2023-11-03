package main

import (
	"log/slog"
	"os"

	"github.com/koha90/podkrepizza-api-v1/config"
	"github.com/koha90/podkrepizza-api-v1/internal/app"
	"github.com/koha90/podkrepizza-api-v1/pkg/logger/sl"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustConfig()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	log.Info("initializing server", slog.String("address", cfg.HTTP.Port))
	log.Debug("logger debug mode enabled")

	if err := app.Run(*cfg, log); err != nil {
		log.Error("error in start application", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
