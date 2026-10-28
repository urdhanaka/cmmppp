package routes

import (
	"cmmppp/internal/api/v1/handlers"
	"cmmppp/internal/config"
	"cmmppp/internal/repositories"
	"cmmppp/internal/services"

	"github.com/go-chi/chi/v5"
)

func registerAuthRoutes(router chi.Router, deps config.Deps) {
	userRepository := repositories.NewUserRepository(deps.DB)
	s := services.NewAuthService(userRepository, deps.TokenManager)
	h := handlers.NewAuthHandler(s)

	router.Post("/login", h.Login)
}
