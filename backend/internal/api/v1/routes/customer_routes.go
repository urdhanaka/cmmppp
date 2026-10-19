package routes

import (
	"cmmppp/internal/api/v1/handler"
	"cmmppp/internal/config"
	"cmmppp/internal/repositories"
	"cmmppp/internal/service"

	"github.com/go-chi/chi/v5"
)

func registerCustomerRoutes(router chi.Router, deps config.Deps) {
	repo := repositories.NewCustomerRepository(deps.DB)
	svc := service.NewCustomerService(repo)
	h := handler.NewCustomerHandler(svc)

	router.Get("/customers", h.GetCustomers)
	router.Post("/customers", h.CreateCustomer)
}
