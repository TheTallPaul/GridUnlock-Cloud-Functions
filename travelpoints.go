package matchmaker

import (
	"context"
	"fmt"
	"log"

	"googlemaps.github.io/maps"
)

func getTravelpoints(startLat, startLng, endLat, endLng float64, apiLey string, maxNodes int) {
	startLocation := fmt.Sprintf("%f", startLat) + "," + fmt.Sprintf("%f", startLng)
	endLocation := fmt.Sprintf("%f", endLat) + "," + fmt.Sprintf("%f", endLng)

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      startLocation,
		Destination: endLocation,
	}

}
