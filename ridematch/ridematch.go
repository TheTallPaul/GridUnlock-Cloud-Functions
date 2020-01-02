// Copyright (c) 2019 Cascade October LLC.

// Package ridematch provides functions that take lists of driver routes and
// rider destinations
package ridematch

import (
	"gonum.org/v1/gonum/stat/distuv"
	"pault.ag/go/haversine"
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
func Matchmake(context context.Context, event FirestoreEvent) error {
	riderRoutes := getUnassignedRiderRoutes()
	driverRoutes := getUnfullDriverRoutes()

	assignedRoutes := matchRidersDrivers(riderRoutes, driverRoutes, 3)
	updateFirebase(assignedRoutes)
}

// matchRidersDrivers finds the shortest distances from a driver route to each
// rider and then weighted randomizes the findings to return a list of matches
func matchRidersDrivers(riderRoutes, driverRoutes,
	timesToRepeat int) []driverRoute {
	var potentialPickups possiblePickups[]

	for _, driverRoute := driverRoutes {
		potentialPickups = append(
			potentialPickups,
			riderDriverClosestDistances(driverRoute, riderRoutes),
		)
	}

	return randomMatch(potentialPickups)

}

// riderDriverClosestDistances finds the closest Haversine distance to a node on
// a drivers route to a rider
func riderDriverClosestDistances(driverRoute driverRoute,
	riderRoutes) float64 {


}

// haversineDistance finds the Haversine distance in meters between two coords
func haversineDistance(coordA, coordB coord) float64 {
	pointA = haversine.Point{Lat: coordA.lat, Lon: coordA.lon}
	pointB = haversine.Point{Lat: coordB.lat, Lon: coordB.lon}

	return float64(pointA.MetresTo(pointB))
}

// randomMatch performs a weighted random match selection, where closer
// distances are weighted heavier than further distances
func randomMatch(potentialPickups []possiblePickups) []riderDriverMatch {
	var matchedDrivers string[]
	var matchedRiders string[]
	var matches riderDriverMatch[]

	for _, pickup := potentialPickups {


	}


}

// weightedInverseRandRider creates a random inverse set and selects one
func weightedInverseRandRider(pickupSet possiblePickups) riderDriverMatch {
	var weightedDistance = inverseArray(
		convertPickupsToArray(distances []riderDistance)
	)
	var driverIndex;

	return riderDriverMatch{
		pickupSet.riderID,
		pickupSet.networkRiderDistances[driverIndex].
	}
}

// inverseArray inverses each element of an array
func inverseArray(array float64[]) float64[] {
	var inverseArray [len(array)]float64

	// Invert each element
	for index, element := array {
		inverseArray[index] = 1.0 / element
	}

	return inverseArray
}


// convertPickupsToArray converts the networkRiderDistances slice in a
// possiblePickups into a weighted indexed array
func convertPickupsToArray(distances []riderDistance) float64[] {
	var allDistances [len(distances)]float64
	for i, riderDistance := distances {
		allDistances[i] = riderDistance.distance
	}

	return allDistances
}
