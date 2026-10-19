package main

import (
	"log"

	"cmmppp/internal/config"
	"cmmppp/internal/server"
	"cmmppp/internal/store"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("ERR: Failed to load config: %v", err)
	}

	db, err := store.NewDB(*cfg)
	if err != nil {
		log.Fatalf("ERR: Failed to connect to database: %v", err)
	}

	server.Run(config.Deps{DB: db})
}
