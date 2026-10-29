package server

import (
	"fmt"
	"net/http"

	"cmmppp/internal/api/v1/routes"
	"cmmppp/internal/config"

	"github.com/go-chi/chi/v5"
)

func newRouter(deps config.Deps) *chi.Mux {
	r := chi.NewRouter()

	r.Use(corsMiddleware)
	r.Use(loggingMiddleware)

	routes.RegisterRoutes(r, deps)

	return r
}

func Run(deps config.Deps) {
	router := newRouter(deps)

	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println(err)
	}
}
