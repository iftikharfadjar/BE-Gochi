package main

import (
	"context"
	"log"
	"time"

	"github.com/iftikharfadjar/BE-Gochi/router"
	_route "github.com/iftikharfadjar/BE-Gochi/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	httpRouter = router.NewChiRouter()
)

func main() {
	const port string = ":3000"

	// serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	// clientOptions := options.Client().
	// 	ApplyURI("mongodb+srv://admin:UXNIhGeK3Rp655Oi@fainacluster.bu653.mongodb.net/coursekita_db").
	// 	SetServerAPIOptions(serverAPIOptions)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, clientOptions)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://admin:UXNIhGeK3Rp655Oi@fainacluster.bu653.mongodb.net/coursekita_db"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)




	_route.InitRoute(httpRouter)

	


	httpRouter.SERVE(port)
}
