package server

import (
	"net/http"

	"cmmppp/internal/api/v1/routes"
	"cmmppp/internal/config"

	"github.com/go-chi/chi/v5"
)

func newRouter(deps config.Deps) *chi.Mux {
	r := chi.NewRouter()

	routes.RegisterRoutes(r, deps)

	return r
}

func Run(deps config.Deps) {
	router := newRouter(deps)

	http.ListenAndServe(":8000", router)
}
