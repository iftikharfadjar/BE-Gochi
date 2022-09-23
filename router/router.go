package router

import (
	"github.com/go-chi/chi/v5"
)

type IRouter interface {
	SERVE(port string)
	GetRouter() *chi.Mux
}
