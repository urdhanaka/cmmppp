package main

import (
	"log/slog"
	"os"

	"cmmppp/internal/config"
	"cmmppp/internal/logger"
	"cmmppp/internal/server"
	"cmmppp/internal/store"
)

func main() {
	slog.SetDefault(logger.New())

	cfg, err := config.New()
	if err != nil {
		slog.Error("failed to load config", "err", err)
		os.Exit(1)
	}

	db, err := store.NewDB(*cfg)
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}

	server.Run(config.Deps{DB: db})
}
