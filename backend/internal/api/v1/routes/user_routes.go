package routes

import (
	"cmmppp/internal/api/v1/handlers"
	"cmmppp/internal/config"
	"cmmppp/internal/repositories"
	"cmmppp/internal/services"

	"github.com/go-chi/chi/v5"
)

func registerUserRoutes(router chi.Router, deps config.Deps) {
	userRepository := repositories.NewUserRepository(deps.DB)
	userService := services.NewUserService(userRepository)
	h := handlers.NewUserHandler(userService)

	router.Get("/users", h.GetUsers)
	router.Post("/user", h.PostUser)
	router.Delete("/user/{userID}", h.DeleteUser)
}
