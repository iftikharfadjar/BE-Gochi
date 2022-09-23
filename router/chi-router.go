package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type chiRouter struct {}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() IRouter {
	return &chiRouter{}
}

func (*chiRouter) SERVE(port string){
	fmt.Printf("chi HTTP Server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher )
}

func (*chiRouter) GetRouter() *chi.Mux {
	return chiDispatcher
}