package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("You found me!"))
	})

	return r
}

func Run() {
	router := NewRouter()

	http.ListenAndServe(":8000", router)
}
