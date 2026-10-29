package main

import (
	"log/slog"
	"os"

	"cmmppp/internal/config"
	"cmmppp/internal/logger"
	"cmmppp/internal/server"
)

func main() {
	slog.SetDefault(logger.New())

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	db, err := config.NewDB(*cfg)
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	tokenManager := config.NewTokenManager(*cfg)

	server.Run(config.Deps{
		DB:           db,
		TokenManager: tokenManager,
	})
}
