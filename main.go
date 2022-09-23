package main

import (
	"fmt"
	"github.com/iftikharfadjar/Be-GoChi/router"
)

var (
	httpRouter = router.NewChiRouter()
)

func main() {
	fmt.Println("Hello, World!")

	const port string = ":3000"
	httpRouter.SERVE(port)
}
