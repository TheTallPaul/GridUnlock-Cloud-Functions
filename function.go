// Copyright (c) 2019 Cascade October LLC.

// Module gridunlockridematch provides packages that can query Firebase for
// unmatched riders and drivers for GridUnlock.
package gridunlockridematch

import (
	"cascadeoctober.com/gridunlockridematch/internal/firebaserepo"
	"cascadeoctober.com/gridunlockridematch/internal/mapsrepo"
	"cascadeoctober.com/gridunlockridematch/ridematch"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"

	"log"
	"net/http"
	"time"
)

// findRoute is a go routine that makes API calls to the map repo and returns
// them as DriverRoutes
func findRoute(drives []firebaserepo.Drive, ch chan ridematch.DriverRoute) {
	for _, drive := range drives {
		ch <- ridematch.DriverRoute{
			Route: mapsrepo.Route(drive.Origin,
				drive.Destination),
			DriverID: drive.DriverID,
		}
	}
	close(ch)
}

// Matchmake queries the database for available rides, matches riders to
// drivers, and updates the database accordingly
func Matchmake(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()

	// Collect rides from collection
	rides := make([]firebaserepo.Ride, 0)
	firebaserepo.FetchCollection(
		ctx,
		"rides",
		func(doc *firestore.DocumentSnapshot) {
			var ride firebaserepo.Ride
			if err := doc.DataTo(&ride); err != nil {
				log.Fatalf(
					"Failed to convert ride json data to "+
						"struct: %v",
					err,
				)
			} else if ride.Active {
				ride.ID = doc.Ref.ID
				rides = append(rides, ride)
			}
		},
	)
	log.Println("Collected rides:", rides)

	// Collect drives from collection
	drives := make([]firebaserepo.Drive, 0)
	firebaserepo.FetchCollection(
		ctx,
		"drives",
		// TO-DO move findRoute() into the loop to optimize API calls
		func(doc *firestore.DocumentSnapshot) {
			var drive firebaserepo.Drive
			if err := doc.DataTo(&drive); err != nil {
				log.Fatalf(
					"Failed to convert drive json data "+
						"to struct: %v",
					err,
				)
			} else {
				drive.ID = doc.Ref.ID
				drives = append(drives, drive)
			}
		},
	)
	log.Println("Collected drives:", drives)

	// Find the driver routes
	ch := make(chan ridematch.DriverRoute)
	go findRoute(drives, ch)
	var driverRoutes []ridematch.DriverRoute
	for route := range ch {
		driverRoutes = append(driverRoutes, route)
	}
	log.Println("Found driver routes:", driverRoutes)

	// Match riders to drivers
	matches := ridematch.MatchRidersDrivers(
		rides,
		driverRoutes,
		1,
		time.Now().UnixNano(),
	)
	log.Println("Matched riders with drivers:", matches)

	firebaserepo.UpdateMatches(ctx, rides, matches)
}
