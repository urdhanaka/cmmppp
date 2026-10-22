package main

import (
	"cmmppp/internal/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("ERR: Failed to load config: %v", err)
	}

	println(cfg.Database.Host)
}
