package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitRoute(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi golang"))
	})
}
