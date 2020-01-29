package ridematch

import (
	"testing"
	"reflect"

	"googlemaps.github.io/maps"
	"google.golang.org/genproto/googleapis/type/latlng"

	"gridunlockridematch/internal/firebaserepo"
)

var bostonRider    = latlng.LatLng{Latitude: 42.4, Longitude: -71.1}
var newYorkRider   = latlng.LatLng{Latitude: 40.7, Longitude: -74.0}
var baltimoreRider = latlng.LatLng{Latitude: 39.3, Longitude: -76.6}
var tacomaRider    = latlng.LatLng{Latitude: 47.2, Longitude: -122.4}
var portlandRider  = latlng.LatLng{Latitude: 45.5, Longitude: -122.6}
var omahaRider     = latlng.LatLng{Latitude: 41.3, Longitude: -95.9}
var lincolnRider   = latlng.LatLng{Latitude: 40.8, Longitude: -96.7}

var MatchRidersDriversTestCases = []struct {
	riderPoints     []firebaserepo.Ride
	driverPaths     []DriverRoute
	timesToRepeat   int
	randSeed        int64
	expectedMatches map[string]string
}{
	{
		[]firebaserepo.Ride{
			firebaserepo.Ride {
				RiderID: "rider1",
				// New York
				PickupLocation: &newYorkRider,
				// Baltimore
				DropoffLocation: &baltimoreRider,
			},
			firebaserepo.Ride {
				RiderID: "rider2",
				// Tacoma
				PickupLocation: &tacomaRider,
				// Portland
				DropoffLocation: &portlandRider,
			},
			firebaserepo.Ride {
				RiderID: "rider3",
				// Omaha
				PickupLocation: &omahaRider,
				// Lincoln
				DropoffLocation:&lincolnRider,
			},
		},
		[]DriverRoute{
			DriverRoute {
				DriverID: "driver1",
				Route: []maps.LatLng{
					// Boston
					maps.LatLng{
						Lat: 42.3602534,
						Lng: -71.0582912,
					},
					// Philadelphia
					maps.LatLng{
						Lat: 39.953346252441406,
						Lng: -75.1633529663086,
					},
					// New York
					maps.LatLng{
						Lat: 40.7127281,
						Lng: -74.0060152,
					},
					// Baltimore
					maps.LatLng{
						Lat: 39.2908816,
						Lng: -76.610759,
					},
					// Washington
					maps.LatLng{
						Lat: 38.895530700683594,
						Lng: -77.0319595336914,
					},
				},
			},
			DriverRoute {
				DriverID: "driver2",
				Route: []maps.LatLng{
					// Seattle
					maps.LatLng{
						Lat: 47.6038321,
						Lng: -122.3300624,
					},
					// Tacoma
					maps.LatLng{
						Lat: 47.2495798,
						Lng: -122.4398746,
					},
					// Portland
					maps.LatLng{
						Lat: 45.5202471,
						Lng: -122.6741949,
					},
					// Eugene
					maps.LatLng{
						Lat: 44.0505054,
						Lng: -123.0950506,
					},
				},
			},
			DriverRoute {
				DriverID: "driver3",
				Route: []maps.LatLng{
					// Sioux Falls
					maps.LatLng{
						Lat: 43.5472794,
						Lng: -96.7294388,
					},
					// Omaha
					maps.LatLng{
						Lat: 41.2587459,
						Lng: -95.9383758,
					},
					// Lincoln
					maps.LatLng{
						Lat: 40.8088861,
						Lng: -96.7077751,
					},
				},
			},

		},
		1,
		1,
		map[string]string{
			"driver1": "rider1",
			"driver2": "rider2",
			"driver3": "rider3",
		},
	},
	{
		[]firebaserepo.Ride{
			firebaserepo.Ride {
				RiderID: "rider1",
				// New York
				PickupLocation: &newYorkRider,
				// Baltimore
				DropoffLocation: &baltimoreRider,
			},
			firebaserepo.Ride {
				RiderID: "rider2",
				// Tacoma
				PickupLocation: &tacomaRider,
				// Portland
				DropoffLocation: &portlandRider,
			},
			firebaserepo.Ride {
				RiderID: "rider3",
				// Omaha
				PickupLocation: &omahaRider,
				// Lincoln
				DropoffLocation: &lincolnRider,
			},
		},
		[]DriverRoute{
			DriverRoute {
				DriverID: "driver1",
				Route: []maps.LatLng{
					// Boston
					maps.LatLng{
						Lat: 42.3602534,
						Lng: -71.0582912,
					},
					// Philadelphia
					maps.LatLng{
						Lat: 39.953346252441406,
						Lng: -75.1633529663086,
					},
					// New York
					maps.LatLng{
						Lat: 40.7127281,
						Lng: -74.0060152,
					},
					// Baltimore
					maps.LatLng{
						Lat: 39.2908816,
						Lng: -76.610759,
					},
					// Washington
					maps.LatLng{
						Lat: 38.895530700683594,
						Lng: -77.0319595336914,
					},
				},
			},
			DriverRoute {
				DriverID: "driver2",
				Route: []maps.LatLng{
					// Seattle
					maps.LatLng{
						Lat: 47.6038321,
						Lng: -122.3300624,
					},
					// Tacoma
					maps.LatLng{
						Lat: 47.2495798,
						Lng: -122.4398746,
					},
					// Portland
					maps.LatLng{
						Lat: 45.5202471,
						Lng: -122.6741949,
					},
					// Eugene
					maps.LatLng{
						Lat: 44.0505054,
						Lng: -123.0950506,
					},
				},
			},
		},
		1,
		1,
		map[string]string{
			"driver1": "rider1",
			"driver2": "rider2",
		},
	},
	{
		[]firebaserepo.Ride{
			firebaserepo.Ride {
				RiderID: "rider1",
				// New York
				PickupLocation: &newYorkRider,
				// Baltimore
				DropoffLocation: &baltimoreRider,
			},
			firebaserepo.Ride {
				RiderID: "rider2",
				// Tacoma
				PickupLocation: &tacomaRider,
				// Portland
				DropoffLocation: &portlandRider,
			},
			firebaserepo.Ride {
				RiderID: "rider3",
				// Omaha
				PickupLocation: &omahaRider,
				// Lincoln
				DropoffLocation: &lincolnRider,
			},
		},
		[]DriverRoute{},
		1,
		1,
		map[string]string{},
	},
}

func TestMatchRidersDrivers(t *testing.T) {
	for _, input := range MatchRidersDriversTestCases {
		matches := MatchRidersDrivers(
			input.riderPoints,
			input.driverPaths,
			input.timesToRepeat,
			input.randSeed,
		)

		if !reflect.DeepEqual(input.expectedMatches, matches) {
			t.Errorf(
				"FAIL: Want matches from riders %v and " +
				" drivers %v to be: %v but we got %v",
				input.riderPoints,
				input.driverPaths,
				input.expectedMatches,
				matches,
			)
		}
	}
}

var randomMatchTestCases = []struct {
	potentialPickups []possiblePickups
	randSeed         int64
	expectedMatches  map[string]string
}{
	{
		[]possiblePickups{
			possiblePickups{
				riderID: "rider1",
				networkRiderDistances: map[string]float64 {
					"driver1": 1.1,
					"driver2": 304.0,
					"driver3": 43.0,
				},
			},
			possiblePickups{
				riderID: "rider2",
				networkRiderDistances: map[string]float64 {
					"driver1": 430.1,
					"driver2": 3.0,
					"driver3": 43.0,
				},
			},
			possiblePickups{
				riderID: "rider3",
				networkRiderDistances: map[string]float64 {
					"driver1": 430.1,
					"driver2": 33.0,
					"driver3": 1.0,
				},
			},
		},
		1,
		map[string]string{
			"driver1": "rider1",
			"driver2": "rider2",
			"driver3": "rider3",
		},
	},
	{
		[]possiblePickups{
			possiblePickups{
				riderID: "rider1",
				networkRiderDistances: map[string]float64 {
					"driver1": 1.1,
				},
			},
			possiblePickups{
				riderID: "rider2",
				networkRiderDistances: map[string]float64 {
					"driver1": 430.1,
				},
			},
			possiblePickups{
				riderID: "rider3",
				networkRiderDistances: map[string]float64 {
					"driver1": 10.1,
					"driver3": 430.1,
				},
			},
		},
		1,
		map[string]string{
			"driver1": "rider1",
			"driver3": "rider3",
		},
	},
	{
		[]possiblePickups{
			possiblePickups{
				riderID: "rider1",
				networkRiderDistances: map[string]float64 {},
			},
			possiblePickups{
				riderID: "rider2",
				networkRiderDistances: map[string]float64 {},
			},
		},
		1,
		map[string]string{},
	},

}

func TestRandomMatch(t *testing.T) {
	for _, input := range randomMatchTestCases {
		matches := randomMatch(input.potentialPickups, input.randSeed)

		if !reflect.DeepEqual(input.expectedMatches, matches) {
			t.Errorf(
				"FAIL: Want matches from %v to be: %v but we " +
				"got %v",
				input.potentialPickups,
				input.expectedMatches,
				matches,
			)
		}
	}
}

var riderDriverClosestDistancesTestCases = []struct {
	DriverRoute      DriverRoute
	riderRoute       firebaserepo.Ride
	expectedDistance float64
	expectedOrder    bool
}{
	{
		DriverRoute {
			DriverID: "driver1",
			Route: []maps.LatLng{
				// Boston
				maps.LatLng{
					Lat: 42.3602534,
					Lng: -71.0582912,
				},
				// Philadelphia
				maps.LatLng{
					Lat: 39.953346252441406,
					Lng: -75.1633529663086,
				},
				// New York
				maps.LatLng{
					Lat: 40.7127281,
					Lng: -74.0060152,
				},
				// Baltimore
				maps.LatLng{
					Lat: 39.2908816,
					Lng: -76.610759,
				},
				// Washington
				maps.LatLng{
					Lat: 38.895530700683594,
					Lng: -77.0319595336914,
				},
			},
		},
		firebaserepo.Ride {
			RiderID: "rider1",
			// New York
			PickupLocation: &newYorkRider,
			// Baltimore
			DropoffLocation: &baltimoreRider,
		},
		2876,
		true,
	},
	{
		DriverRoute {
			DriverID: "driver1",
			Route: []maps.LatLng{
				// Boston
				maps.LatLng{
					Lat: 42.3602534,
					Lng: -71.0582912,
				},
				// Philadelphia
				maps.LatLng{
					Lat: 39.953346252441406,
					Lng: -75.1633529663086,
				},
				// New York
				maps.LatLng{
					Lat: 40.7127281,
					Lng: -74.0060152,
				},
				// Baltimore
				maps.LatLng{
					Lat: 39.2908816,
					Lng: -76.610759,
				},
				// Washington
				maps.LatLng{
					Lat: 38.895530700683594,
					Lng: -77.0319595336914,
				},
			},
		},
		firebaserepo.Ride {
			RiderID: "rider1",
			// Baltimore
			PickupLocation: &baltimoreRider,
			// New York
			DropoffLocation: &newYorkRider,
		},
		2876,
		false,
	},
	{
		DriverRoute {
			DriverID: "driver1",
			Route: []maps.LatLng{
				// Boston
				maps.LatLng{
					Lat: 42.3602534,
					Lng: -71.0582912,
				},
				// Philadelphia
				maps.LatLng{
					Lat: 39.953346252441406,
					Lng: -75.1633529663086,
				},
				// Washington
				maps.LatLng{
					Lat: 38.895530700683594,
					Lng: -77.0319595336914,
				},
			},
		},
		firebaserepo.Ride {
			RiderID: "rider1",
			// New York
			PickupLocation: &newYorkRider,
			// Baltimore
			DropoffLocation: &baltimoreRider,
		},
		187325,
		true,
	},
	{
		DriverRoute {
			DriverID: "driver1",
			Route: []maps.LatLng{
				// Boston
				maps.LatLng{
					Lat: 42.3602534,
					Lng: -71.0582912,
				},
				// Philadelphia
				maps.LatLng{
					Lat: 39.953346252441406,
					Lng: -75.1633529663086,
				},
				// Washington
				maps.LatLng{
					Lat: 38.895530700683594,
					Lng: -77.0319595336914,
				},
			},
		},
		firebaserepo.Ride {
			RiderID: "rider1",
			// Baltimore
			PickupLocation: &baltimoreRider,
			// New York
			DropoffLocation: &newYorkRider,
		},
		187325,
		false,
	},
	{
		DriverRoute {
			DriverID: "driver1",
			Route: []maps.LatLng{
				// Boston
				maps.LatLng{
					Lat: 42.3602534,
					Lng: -71.0582912,
				},
			},
		},
		firebaserepo.Ride {
			RiderID: "rider1",
			// New York
			PickupLocation: &newYorkRider,
			// Baltimore
			DropoffLocation: &baltimoreRider,
		},
		883650,
		true,
	},
	{
		DriverRoute {
			DriverID: "driver1",
			Route:	[]maps.LatLng{
				// Boston
				maps.LatLng{
					Lat: 42.3602534,
					Lng: -71.0582912,
				},
			},
		},
		firebaserepo.Ride {
			RiderID: "rider1",
			// Boston
			PickupLocation: &bostonRider,
			// Boston
			DropoffLocation: &bostonRider,
		},
		11184,
		true,
	},

}

func TestRiderDriverClosestDistances(t *testing.T) {
	for _, input := range riderDriverClosestDistancesTestCases {
		distance, order := riderDriverClosestDistances(
			input.DriverRoute, input.riderRoute)

		if input.expectedDistance != distance {
			t.Errorf(
				"FAIL: Want closest total distance for Route " +
				"%v picking up %v to be: %v but we got %v",
				input.DriverRoute,
				input.riderRoute,
				input.expectedDistance,
				distance,
			)
		}

		if input.expectedOrder != order {
			t.Errorf(
				"FAIL: Want correct ordering for Route %v " +
				"picking up %v to be: %v but we got %v",
				input.DriverRoute,
				input.riderRoute,
				input.expectedOrder,
				order,
			)
		}
	}
}

var haversineDistanceTestCases = []struct {
	coordA         maps.LatLng
	coordB         maps.LatLng
	expectedMeters float64
}{
	{
		maps.LatLng{Lat: 38.89768, Lng: -77.03653},
		maps.LatLng{Lat: 38.89736, Lng: -77.04173},
		451,
	},
	{
		maps.LatLng{Lat: 51.510357, Lng: -0.116773},
		maps.LatLng{Lat: 38.889931, Lng: -77.009003},
		5897658,
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range haversineDistanceTestCases {
		meters := haversineDistance(input.coordA, input.coordB)

		if input.expectedMeters != meters {
			t.Errorf(
				"FAIL: Want distance from %v to %v to be: %v " +
				"but we got %v",
				input.coordA,
				input.coordB,
				input.expectedMeters,
				meters,
			)
		}
	}
}

var weightedInverseRandRiderTestCases = []struct {
	pickupSet   possiblePickups
	prevMatches map[string]string
	randSeed    int64
	expectedID  string
	expectedOK  bool
}{
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: map[string]float64 {
				"driver1": 1.0,
				"driver2": 100.0,
			},
		},
		map[string]string {},
		1,
		"driver1",
		true,
	},
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: map[string]float64 {
				"driver1": 430.1,
				"driver2": 3.0,
				"driver3": 43.0,
			},
		},
		map[string]string {},
		1,
		"driver2",
		true,
	},
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: map[string]float64 {
				"driver1": 430.1,
				"driver2": 3.0,
				"driver3": 43.0,
			},
		},
		map[string]string {"driver2": "rider2"},
		1,
		"driver3",
		true,
	},
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: map[string]float64 {
				"driver1": 430.1,
				"driver2": 3.0,
				"driver3": 43.0,
			},
		},
		map[string]string {
			"driver1": "rider2",
			"driver2": "rider3",
			"driver3": "rider4",
		},
		1,
		"",
		false,
	},
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: map[string]float64 {},
		},
		map[string]string {},
		1,
		"",
		false,
	},
}


func TestWeightedInverseRandRider(t *testing.T) {
	for _, input := range weightedInverseRandRiderTestCases {
		DriverID, matched := weightedInverseRandRider(
			input.pickupSet,
			input.prevMatches,
			input.randSeed,
		)

		if input.expectedID != DriverID || input.expectedOK != matched {
			t.Errorf(
				"FAIL: Want match to be: (%v, %v) but we got " +
				"(%v, %v)",
				input.expectedID,
				input.expectedOK,
				DriverID,
				matched,
			)
		}
	}
}
