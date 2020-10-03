package main

import (
	"context"
	"github.com/sunil206b/jwt_api/src/app"
	"github.com/sunil206b/jwt_api/src/driver"
	"log"
)

func main() {
	client, err := driver.MongoDB()
	ctx := context.Background()
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("error while disconnecting from mongodb %v\n", err)
		}
	}()
	collection := client.Database("userdb").Collection("user")
	if err != nil {
		log.Fatalln(err)
	}
	srv := app.NewRouter(collection, ctx)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v\n", err)
	}
}
