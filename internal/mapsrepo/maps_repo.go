package mapsrepo

import (
	"context"
	"log"
	"fmt"

	"googlemaps.github.io/maps"
	"google.golang.org/genproto/googleapis/type/latlng"
)

// client is a Google Maps client, reused between function invocations.
var client *maps.Client

func init(){
	var err error
	client, err = maps.NewClient(maps.WithAPIKey(mapsAPIKey))
	if err != nil {
		log.Fatalf("maps.NewClient: %s", err)
	}
}

// Route makes Google Maps Direction calls between an origin and destination and
// returns an array of all the coordinates where a turn is required
func Route(origin, destination *latlng.LatLng) []maps.LatLng {
	request := &maps.DirectionsRequest{
		Origin: fmt.Sprintf("%f", origin.Latitude) + "," +
			fmt.Sprintf("%f", origin.Longitude),
		Destination: fmt.Sprintf("%f", destination.Latitude) + "," +
			fmt.Sprintf("%f", destination.Longitude),
	}
	route, _, err := client.Directions(context.Background(), request)
	if err != nil {
		log.Fatalf("Directions fatal error: %s", err)
	}

	if len(route) < 1 {
		log.Fatalf(
			"Directions error: Can't find route between %v,%v " +
			"and %v,%v",
			origin.Latitude,
			origin.Longitude,
			destination.Latitude,
			destination.Longitude,
		)

		return []maps.LatLng{}
	}

	driverRoute := []maps.LatLng{route[0].Legs[0].Steps[0].StartLocation}

	// TO-DO, add fake nodes for long distances without turns (<100 total)
	for _, step := range route[0].Legs[0].Steps {
		driverRoute = append(driverRoute, step.EndLocation)
	}

	return driverRoute
}
