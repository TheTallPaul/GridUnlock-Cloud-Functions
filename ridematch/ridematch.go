// Copyright (c) 2019 Cascade October LLC.

// Package ridematch provides functions that take lists of driver routes and
// rider destinations
package ridematch

import (
	"pault.ag/go/haversine"
	"math/rand"
	// "time"
)

// A riderDistance is the distance from a rider to a driver's node point
type riderDistance struct {
	driverID string
	distance float64
}


// A possiblePickups is a list of all of the potential drivers that can pick a
// rider up
type possiblePickups struct {
	riderID string
	networkRiderDistances []riderDistance
}

// A riderDriverMatch is a pairing of a driver and a rider
type riderDriverMatch struct {
	riderID string
	driverID string
}

// matchmake makes calls for the unassigned riders and drivers, then passes them
// to other functions to make
// func Matchmake(context context.Context, event FirestoreEvent) error {
// 	riderRoutes := getUnassignedRiderRoutes()
// 	driverRoutes := getUnfullDriverRoutes()

// 	assignedRoutes := matchRidersDrivers(riderRoutes, driverRoutes, 3)
// 	updateFirebase(assignedRoutes)
// }

// matchRidersDrivers finds the shortest distances from a driver route to each
// rider and then weighted randomizes the findings to return a list of matches
// func matchRidersDrivers(riderRoutes, driverRoutes,
// 	timesToRepeat int) []driverRoute {
// 	var potentialPickups []possiblePickups

// 	for _, driverRoute := driverRoutes {
// 		potentialPickups = append(
// 			potentialPickups,
// 			riderDriverClosestDistances(driverRoute, riderRoutes),
// 		)
// 	}

// 	return randomMatch(potentialPickups)

// }

// riderDriverClosestDistances finds the closest Haversine distance to a node on
// a drivers route to a rider
// func riderDriverClosestDistances(driverRoute driverRoute,
// 	riderRoutes) float64 {


// }

// haversineDistance finds the Haversine distance in meters between two coords
func haversineDistance(coordA, coordB coord) float64 {
	var pointA = haversine.Point{Lat: coordA.lat, Lon: coordA.lng}
	var pointB = haversine.Point{Lat: coordB.lat, Lon: coordB.lng}

	return float64(pointA.MetresTo(pointB))
}

// randomMatch performs a weighted random match selection, where closer
// distances are weighted heavier than further distances
// func randomMatch(potentialPickups []possiblePickups) []riderDriverMatch {
// 	var matchedDrivers []string
// 	var matchedRiders []string
// 	var matches []riderDriverMatch

// 	for _, pickup := range potentialPickups {

// 	}

// 	return 0
// }


// weightedInverseRandRider creates a random inverse set and selects one rider /
// driver pairing
func weightedInverseRandRider(pickupSet possiblePickups,
	randSeed int64) riderDriverMatch {
	weightedDistance := inverseSlice(
		convertPickupsToSlice(pickupSet.networkRiderDistances))
	distancesTotal := sumSlice(weightedDistance)
	totalWeight := 0.0
	winnerIndex := len(weightedDistance) - 1

	rand.Seed(randSeed)
	targetWeight := rand.Float64() * distancesTotal

	// Adds each distance until the random target is reached (weighted
	// random selection)
	for index, weight := range weightedDistance {
		totalWeight += weight
		if (totalWeight >= targetWeight) {
			winnerIndex = index
			break
		}
	}

	return riderDriverMatch{
		pickupSet.riderID,
		pickupSet.networkRiderDistances[winnerIndex].driverID,
	}
}

// inverseSlice inverses each element of an slice
func inverseSlice(slice []float64) []float64 {
	var inverseSlice []float64

	// Invert each element
	for _, element := range slice {
		inverseSlice = append(inverseSlice, 1 / element)
	}

	return inverseSlice
}

// sumSlice returns the sum of a slice
func sumSlice(slice []float64) float64 {
	total := 0.0

	for _, x := range slice {
		total += x
	}

	return total
}


// convertPickupsToSlice converts the networkRiderDistances slice in a
// possiblePickups into a weighted slice
func convertPickupsToSlice(distances []riderDistance) []float64 {
	var allDistances []float64
	for _, riderDistance := range distances {
		allDistances = append(allDistances, riderDistance.distance)
	}

	return allDistances
}
