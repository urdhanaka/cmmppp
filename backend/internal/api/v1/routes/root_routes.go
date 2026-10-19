package routes

import (
	"net/http"

	"cmmppp/internal/config"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, deps config.Deps) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You found me!"))
	})

	router.Route("/api/v1", func(r chi.Router) {
		registerUserRoutes(r, deps)
		registerCustomerRoutes(r, deps)
	})
}
