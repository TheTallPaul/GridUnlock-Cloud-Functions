package ridematch

import (
	"context"
	"fmt"
	"log"

	"./structs"

	"googlemaps.github.io/maps"
)

func travelpoints(startCoord, endCoord coord, apiKey string, maxNodes int) {
	startLocation := fmt.Sprintf("%f", startCoord.lat) + "," +
		fmt.Sprintf("%f", startCoord.lng)
	endLocation := fmt.Sprintf("%f", endCoord.lat) + "," +
		fmt.Sprintf("%f", endCoord.lng)

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      startLocation,
		Destination: endLocation,
	}

}
