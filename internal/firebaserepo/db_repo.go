package firebaserepo

import (
	"log"
	"context"

	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"

	"google.golang.org/api/iterator"
)

// client is a Firestore client, reused between function invocations.
var client *firestore.Client

func init() {
	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx    := context.Background()
	config := &firebase.Config{ProjectID: projectID}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

// FetchCollection returns all documents from the chosen collection and allows a
// function to be performed on the collection
func FetchCollection(ctx context.Context, collectionName string,
	collector func(doc *firestore.DocumentSnapshot)) {
	iter := client.Collection(collectionName).Documents(ctx)

	for {
	        doc, err := iter.Next()
	        if err == iterator.Done {
	                break
	        }
	        if err != nil {
	                log.Fatalf("Failed to iterate: %v", err)
	        } else {
	        	collector(doc)
	        }
	}
}
