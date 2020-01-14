package firebaserepo

import (
	"log"
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"google.golang.org/api/option"
)

var client *db.Client

func init() {
	ctx := context.Background()
	config := &firebase.Config{DatabaseURL: firebaseURL}

	app, error := firebase.NewApp(ctx, config)
	if error != nil {
		log.Fatalf("error initializing app: %v\n", error)
	}

	client, error = app.Database(ctx)
	if error != nil {
		log.Fatalf("error getting database: %v\n", error)
	}
}

func FetchRides(ctx context.Context) error {
	ref := client.NewRef("rides")
	var ride Ride

	if error := ref.Get(ctx, &ride); error != nil {
		log.Fatalln("Error reading ride:", error)
	}

	fmt.Println(ref.Path)
}
