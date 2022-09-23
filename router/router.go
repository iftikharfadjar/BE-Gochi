package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type IRouter interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
	GetRouter() chi.Router
}
