package firebaserepo

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/iterator"
)

// client is a Firestore client, reused between function invocations.
var client *firestore.Client

func init() {
	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx := context.Background()
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

// FetchUser returns a Firestore user document matching the provided ID
func FetchUser(ctx context.Context, userID string) (User, error) {
	var user User
	doc, err := client.Doc("users/" + userID).Get(ctx)
	if err != nil {
		return user, errors.New("Fetching user " + userID + " failed")
	}
	if err := doc.DataTo(&user); err != nil {
		return user, errors.New("Failed to convert user json data to " +
			"struct")
	}

	return user, nil
}

// UpdateMatches updates the Ride and User documents with matched rides
func UpdateMatches(ctx context.Context, rides []Ride,
	matches map[string]string) {
	// TO-DO account for the possiblity that rider could have multiple
	// active rides due to multiple requests and a failure to mark the
	// previous ones incomplete
	for _, ride := range rides {
		ride.DriverID = matches[ride.RiderID]

		// Fetch the driver document
		driver, err := FetchUser(ctx, ride.DriverID)
		if err != nil {
			log.Println(err)
			continue
		}

		// Modify the driver
		if driver.AcceptedRides == nil {
			driver.AcceptedRides = make(map[string]time.Time)
		}
		driver.AcceptedRides[ride.ID] = time.Now()

		// Add the ride to the driver's accepted rides
		writeResults, err := client.Batch().
			Set(client.Doc("rides/"+ride.ID), ride).
			Set(client.Doc("users/"+ride.DriverID), driver).
			Commit(ctx)

		if err != nil {
			log.Fatalf("client.Batch: %v", err)
		} else {
			log.Println("Wrote to db: ", writeResults)
		}
	}
}
