package routes

import (
	"cmmppp/internal/api/v1/handler"
	"cmmppp/internal/config"
	"cmmppp/internal/repositories"

	"github.com/go-chi/chi/v5"
)

func registerUserRoutes(router chi.Router, deps config.Deps) {
	h := handler.NewUserHandler(repositories.NewUserRepository(deps.DB))

	router.Get("/users", h.GetUsers)
}
