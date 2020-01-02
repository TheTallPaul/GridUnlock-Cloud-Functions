package firebaserepo

import (
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"

	"./const"
)

var client *db.Client

func init() {
	context := context.Background()
	config := &firebase.Config{DatabaseURL: firebaseURL}
	app, error := firebase.NewApp(context, conf)
	if error != nil {
		log.Fatalf("firebase.NewApp: %v", error)
	}
	client, error = app.Database(context)
	if error != nil {
		log.Fatalf("app.Firestore: %v", error)
	}
}
