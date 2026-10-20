package routes

import (
	"net/http"
	"os"

	"cmmppp/internal/config"

	_ "cmmppp/docs/v1"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func RegisterRoutes(router *chi.Mux, deps config.Deps) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You found me!"))
	})

	router.Route("/api/v1", func(r chi.Router) {
		// swagger, only enabled on development
		if os.Getenv("ENV") == "development" {
			r.Get("/swagger/*", httpSwagger.Handler(
				httpSwagger.URL("http://localhost:8000/api/v1/swagger/doc.json"),
			))
		}

		registerUserRoutes(r, deps)
		registerCustomerRoutes(r, deps)
	})
}
