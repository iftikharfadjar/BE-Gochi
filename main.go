package main

import (
	"fmt"
	"github.com/iftikharfadjar/BE-Gochi/router"
	_route "github.com/iftikharfadjar/BE-Gochi/routes"
)

var (
	httpRouter = router.NewChiRouter()
)

func main() {
	fmt.Println("Hello, World!")

	const port string = ":3000"

	_route.InitRoute(httpRouter)
	httpRouter.SERVE(port)
}
