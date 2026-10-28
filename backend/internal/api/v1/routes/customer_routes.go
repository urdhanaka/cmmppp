package routes

import (
	"cmmppp/internal/api/v1/handlers"
	"cmmppp/internal/config"
	"cmmppp/internal/repositories"
	"cmmppp/internal/services"

	"github.com/go-chi/chi/v5"
)

func registerCustomerRoutes(router chi.Router, deps config.Deps) {
	r := repositories.NewCustomerRepository(deps.DB)
	s := services.NewCustomerService(r)
	h := handlers.NewCustomerHandler(s)

	router.Get("/customers", h.GetCustomers)
	router.Post("/customers", h.CreateCustomer)
}
