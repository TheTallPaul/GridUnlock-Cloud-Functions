// Copyright (c) 2019 Cascade October LLC.

// Package ridematch provides functions that take lists of driver Routes and
// rider destinations
package ridematch

import (
	"math"
	"math/rand"

	"googlemaps.github.io/maps"
	"pault.ag/go/haversine"

	"gridunlockridematch/internal/firebaserepo"
)

// A possiblePickups is a list of all of the potential drivers that can pick a
// rider up
type possiblePickups struct {
	riderID               string
	networkRiderDistances map[string]float64
}

// A DriverRoute is the sequential route a driver takes
type DriverRoute struct {
	DriverID string
	Route    []maps.LatLng
}

// MatchRidersDrivers finds the shortest distances from a driver route to each
// rider and then weighted randomizes the findings to return a list of matches
func MatchRidersDrivers(riderPoints []firebaserepo.Ride,
	driverPaths []DriverRoute, timesToRepeat int,
	randSeed int64) map[string]string {
	var potentialPickups []possiblePickups

	for _, riderCoord := range riderPoints {
		riderPickups := possiblePickups {
			riderID: riderCoord.RiderID,
			networkRiderDistances: map[string]float64{},
		}
		for _, driverPath := range driverPaths {
			distance, correctOrder := riderDriverClosestDistances(
				driverPath, riderCoord)
			if correctOrder {
				riderPickups.networkRiderDistances[
					driverPath.DriverID] = distance
			}
		}
		potentialPickups = append(potentialPickups, riderPickups)
	}

	return randomMatch(potentialPickups, randSeed)
}

// riderDriverClosestDistances finds the total closest Haversine distance to a
// node on a drivers route to a rider's starting point and endpoints. Also
// checks whether the directions for both participants match (ex. driver is
// going from SF to LA, but rider is going from LA to SF)
func riderDriverClosestDistances(driverPath DriverRoute,
	riderCoords firebaserepo.Ride) (float64, bool) {
	shortestStartDistance, shortestEndDistance := math.MaxFloat64,
		math.MaxFloat64
	startIndex, endIndex := -1, -1

	for index, driverNode := range driverPath.Route {
		startDistance := haversineDistance(
			// TO-DO Find a way to include the Ride struct without
			// having to do type conversions
			maps.LatLng{
				Lat: riderCoords.PickupLocation.Latitude,
				Lng: riderCoords.PickupLocation.Longitude,
			},
			driverNode,
		)
		endDistance := haversineDistance(
			maps.LatLng{
				Lat: riderCoords.DropoffLocation.Latitude,
				Lng: riderCoords.DropoffLocation.Longitude,
			},
			driverNode,
		)

		if startDistance < shortestStartDistance {
			shortestStartDistance = startDistance
			startIndex = index
		}
		if endDistance < shortestEndDistance {
			shortestEndDistance = endDistance
			endIndex = index
		}
	}

	// Check to make sure that the rider's destination doesn't come before
	// the start point
	correctOrder := true
	if startIndex > endIndex {
		correctOrder = false
	}

	return shortestStartDistance + shortestEndDistance, correctOrder
}

// haversineDistance finds the rounded Haversine distance in meters between two
// coords
func haversineDistance(coordA, coordB maps.LatLng) float64 {
	pointA := haversine.Point{Lat: coordA.Lat, Lon: coordA.Lng}
	pointB := haversine.Point{Lat: coordB.Lat, Lon: coordB.Lng}

	return math.Round(float64(pointA.MetresTo(pointB)))
}

// randomMatch performs a weighted random match selection on all riders, where
// closer distances are weighted heavier than further distances and returns it
// as a map of "DriverID": "RiderID"
func randomMatch(potentialPickups []possiblePickups,
	randSeed int64) map[string]string {
	matches := make(map[string]string)

	for _, pickup := range potentialPickups {
		DriverID, matched := weightedInverseRandRider(
			pickup, matches, randSeed)
		if matched {
			matches[DriverID] = pickup.riderID
		}
	}

	return matches
}


// weightedInverseRandRider matches a rider to one of their possible drivers
// using inverse random weighting and returns the matching driver ID
func weightedInverseRandRider(pickupSet possiblePickups,
	alreadyMatched map[string]string, randSeed int64) (string, bool) {
	inverseDistanceTotal, totalNum := 0.0, 0.0
	matchedDriverID                := ""
	rand.Seed(randSeed)

	for DriverID, distance := range pickupSet.networkRiderDistances {
		if _, matched := alreadyMatched[DriverID]; !matched {
			inverseDistanceTotal += 1 / distance
		} else {
			// Clean up previous matches
			delete(pickupSet.networkRiderDistances, DriverID)
		}
	}

	targetNum := rand.Float64() * inverseDistanceTotal

	// Adds each distance until the random target is reached (weighted
	// random selection)
	for DriverID, distance := range pickupSet.networkRiderDistances {
		totalNum += 1 / distance
		if (totalNum >= targetNum) {
			matchedDriverID = DriverID
			break
		}
	}

	if len(matchedDriverID) > 0 {
		return matchedDriverID, true
	}

	return matchedDriverID, false
}
