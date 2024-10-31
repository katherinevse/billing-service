package main

import (
	"log/slog"
	"main/app/internal/config"
	"main/app/pkg/utils"
	"os"
)

var (
	cfg    *config.Config
	logger *slog.Logger
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func init() {
	cfg = utils.LoadConfig("./config/app.yaml")
	logger = setupLogger(cfg.LoggerConfig.Level)

}

func main() {

}

func setupLogger(env string) *slog.Logger {
	switch env {
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return logger
}
