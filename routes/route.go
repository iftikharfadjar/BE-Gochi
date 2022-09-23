package routes

import (
	"github.com/iftikharfadjar/BE-Gochi/router"
	"net/http"
)

func InitRoute(r router.IRouter) {
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi golang"))
	})
}
