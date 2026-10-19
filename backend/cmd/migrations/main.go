package main

import (
	"log"

	"cmmppp/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("ERR: Failed to load config: %v", err)
	}

	println(cfg.Database.Host)
}
