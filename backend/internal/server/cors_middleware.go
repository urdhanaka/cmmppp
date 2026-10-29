package server

import (
	"net/http"

	"github.com/go-chi/cors"
)

func corsMiddleware(next http.Handler) http.Handler {
	c := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "PATCH"},
		AllowedHeaders: []string{
			"Accept", "Authorization", "Content-Type", "X-CSRF-Token",
			"Content-Length", "Accept-Encoding", "Origin", "Cache-Control", "X-Requested-With",
		},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	})

	return c(next)
}
