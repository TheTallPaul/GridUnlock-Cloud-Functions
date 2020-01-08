package ridematch

import (
	"testing"
	"reflect"
)

var matchRidersDriversTestCases = []struct {
	riderPoints     []riderStartEnd
	driverPaths     []driverRoute
	timesToRepeat   int
	randSeed        int64
	expectedMatches map[string]string
}{
	{
		[]riderStartEnd{
			riderStartEnd {
				riderID: "rider1",
				// New York
				start: coord{
					lat: 40.7,
					lng: -74.0,
				},
				// Baltimore
				end: coord{
					lat: 39.3,
					lng: -76.6,
				},
			},
			riderStartEnd {
				riderID: "rider2",
				// Tacoma
				start: coord{
					lat: 47.2,
					lng: -122.4,
				},
				// Portland
				end: coord{
					lat: 45.5,
					lng: -122.6,
				},
			},
			riderStartEnd {
				riderID: "rider3",
				// Omaha
				start: coord{
					lat: 41.3,
					lng: -95.9,
				},
				// Lincoln
				end: coord{
					lat: 40.8,
					lng: -96.7,
				},
			},
		},
		[]driverRoute{
			driverRoute {
				driverID: "driver1",
				route: []coord{
					// Boston
					coord{
						lat: 42.3602534,
						lng: -71.0582912,
					},
					// Philadelphia
					coord{
						lat: 39.953346252441406,
						lng: -75.1633529663086,
					},
					// New York
					coord{
						lat: 40.7127281,
						lng: -74.0060152,
					},
					// Baltimore
					coord{
						lat: 39.2908816,
						lng: -76.610759,
					},
					// Washington
					coord{
						lat: 38.895530700683594,
						lng: -77.0319595336914,
					},
				},
			},
			driverRoute {
				driverID: "driver2",
				route: []coord{
					// Seattle
					coord{
						lat: 47.6038321,
						lng: -122.3300624,
					},
					// Tacoma
					coord{
						lat: 47.2495798,
						lng: -122.4398746,
					},
					// Portland
					coord{
						lat: 45.5202471,
						lng: -122.6741949,
					},
					// Eugene
					coord{
						lat: 44.0505054,
						lng: -123.0950506,
					},
				},
			},
			driverRoute {
				driverID: "driver3",
				route: []coord{
					// Sioux Falls
					coord{
						lat: 43.5472794,
						lng: -96.7294388,
					},
					// Omaha
					coord{
						lat: 41.2587459,
						lng: -95.9383758,
					},
					// Lincoln
					coord{
						lat: 40.8088861,
						lng: -96.7077751,
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
		[]riderStartEnd{
			riderStartEnd {
				riderID: "rider1",
				// New York
				start: coord{
					lat: 40.7,
					lng: -74.0,
				},
				// Baltimore
				end: coord{
					lat: 39.3,
					lng: -76.6,
				},
			},
			riderStartEnd {
				riderID: "rider2",
				// Tacoma
				start: coord{
					lat: 47.2,
					lng: -122.4,
				},
				// Portland
				end: coord{
					lat: 45.5,
					lng: -122.6,
				},
			},
			riderStartEnd {
				riderID: "rider3",
				// Omaha
				start: coord{
					lat: 41.3,
					lng: -95.9,
				},
				// Lincoln
				end: coord{
					lat: 40.8,
					lng: -96.7,
				},
			},
		},
		[]driverRoute{
			driverRoute {
				driverID: "driver1",
				route: []coord{
					// Boston
					coord{
						lat: 42.3602534,
						lng: -71.0582912,
					},
					// Philadelphia
					coord{
						lat: 39.953346252441406,
						lng: -75.1633529663086,
					},
					// New York
					coord{
						lat: 40.7127281,
						lng: -74.0060152,
					},
					// Baltimore
					coord{
						lat: 39.2908816,
						lng: -76.610759,
					},
					// Washington
					coord{
						lat: 38.895530700683594,
						lng: -77.0319595336914,
					},
				},
			},
			driverRoute {
				driverID: "driver2",
				route: []coord{
					// Seattle
					coord{
						lat: 47.6038321,
						lng: -122.3300624,
					},
					// Tacoma
					coord{
						lat: 47.2495798,
						lng: -122.4398746,
					},
					// Portland
					coord{
						lat: 45.5202471,
						lng: -122.6741949,
					},
					// Eugene
					coord{
						lat: 44.0505054,
						lng: -123.0950506,
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
		[]riderStartEnd{
			riderStartEnd {
				riderID: "rider1",
				// New York
				start: coord{
					lat: 40.7,
					lng: -74.0,
				},
				// Baltimore
				end: coord{
					lat: 39.3,
					lng: -76.6,
				},
			},
			riderStartEnd {
				riderID: "rider2",
				// Tacoma
				start: coord{
					lat: 47.2,
					lng: -122.4,
				},
				// Portland
				end: coord{
					lat: 45.5,
					lng: -122.6,
				},
			},
			riderStartEnd {
				riderID: "rider3",
				// Omaha
				start: coord{
					lat: 41.3,
					lng: -95.9,
				},
				// Lincoln
				end: coord{
					lat: 40.8,
					lng: -96.7,
				},
			},
		},
		[]driverRoute{},
		1,
		1,
		map[string]string{},
	},
}

func TestMatchRidersDrivers(t *testing.T) {
	for _, input := range matchRidersDriversTestCases {
		matches := matchRidersDrivers(
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
	driverRoute      driverRoute
	riderRoute       riderStartEnd
	expectedDistance float64
	expectedOrder    bool
}{
	{
		driverRoute {
			driverID: "driver1",
			route: []coord{
				// Boston
				coord{
					lat: 42.3602534,
					lng: -71.0582912,
				},
				// Philadelphia
				coord{
					lat: 39.953346252441406,
					lng: -75.1633529663086,
				},
				// New York
				coord{
					lat: 40.7127281,
					lng: -74.0060152,
				},
				// Baltimore
				coord{
					lat: 39.2908816,
					lng: -76.610759,
				},
				// Washington
				coord{
					lat: 38.895530700683594,
					lng: -77.0319595336914,
				},
			},
		},
		riderStartEnd {
			riderID: "rider1",
			// New York
			start: coord{
				lat: 40.71,
				lng: -74.0,
			},
			// Baltimore
			end: coord{
				lat: 39.3,
				lng: -76.6,
			},
		},
		1964,
		true,
	},
	{
		driverRoute {
			driverID: "driver1",
			route: []coord{
				// Boston
				coord{
					lat: 42.3602534,
					lng: -71.0582912,
				},
				// Philadelphia
				coord{
					lat: 39.953346252441406,
					lng: -75.1633529663086,
				},
				// New York
				coord{
					lat: 40.7127281,
					lng: -74.0060152,
				},
				// Baltimore
				coord{
					lat: 39.2908816,
					lng: -76.610759,
				},
				// Washington
				coord{
					lat: 38.895530700683594,
					lng: -77.0319595336914,
				},
			},
		},
		riderStartEnd {
			riderID: "rider1",

			// Baltimore
			start: coord{
				lat: 39.3,
				lng: -76.6,
			},
			// New York
			end: coord{
				lat: 40.7,
				lng: -74.0,
			},
		},
		2876,
		false,
	},
	{
		driverRoute {
			driverID: "driver1",
			route: []coord{
				// Boston
				coord{
					lat: 42.3602534,
					lng: -71.0582912,
				},
				// Philadelphia
				coord{
					lat: 39.953346252441406,
					lng: -75.1633529663086,
				},
				// Washington
				coord{
					lat: 38.895530700683594,
					lng: -77.0319595336914,
				},
			},
		},
		riderStartEnd {
			riderID: "rider1",
			// New York
			start: coord{
				lat: 40.7,
				lng: -74.0,
			},
			// Baltimore
			end: coord{
				lat: 39.3,
				lng: -76.6,
			},
		},
		187325,
		true,
	},
	{
		driverRoute {
			driverID: "driver1",
			route: []coord{
				// Boston
				coord{
					lat: 42.3602534,
					lng: -71.0582912,
				},
				// Philadelphia
				coord{
					lat: 39.953346252441406,
					lng: -75.1633529663086,
				},
				// Washington
				coord{
					lat: 38.895530700683594,
					lng: -77.0319595336914,
				},
			},
		},
		riderStartEnd {
			riderID: "rider1",
			// Baltimore
			start: coord{
				lat: 39.3,
				lng: -76.6,
			},
			// New York
			end: coord{
				lat: 40.7,
				lng: -74.0,
			},
		},
		187325,
		false,
	},
	{
		driverRoute {
			driverID: "driver1",
			route: []coord{
				// Boston
				coord{
					lat: 42.3602534,
					lng: -71.0582912,
				},
			},
		},
		riderStartEnd {
			riderID: "rider1",
			// New York
			start: coord{
				lat: 40.7,
				lng: -74.0,
			},
			// Baltimore
			end: coord{
				lat: 39.3,
				lng: -76.6,
			},
		},
		883650,
		true,
	},
	{
		driverRoute {
			driverID: "driver1",
			route:	[]coord{
				// Boston
				coord{
					lat: 42.3602534,
					lng: -71.0582912,
				},
			},
		},
		riderStartEnd {
			riderID: "rider1",
			// Boston
			start: coord{
				lat: 42.4,
				lng: -71.1,
			},
			// Boston
			end: coord{
				lat: 42.4,
				lng: -71.1,
			},
		},
		11184,
		true,
	},

}

func TestRiderDriverClosestDistances(t *testing.T) {
	for _, input := range riderDriverClosestDistancesTestCases {
		distance, order := riderDriverClosestDistances(
			input.driverRoute, input.riderRoute)

		if input.expectedDistance != distance {
			t.Errorf(
				"FAIL: Want closest total distance for route " +
				"%v picking up %v to be: %v but we got %v",
				input.driverRoute,
				input.riderRoute,
				input.expectedDistance,
				distance,
			)
		}

		if input.expectedOrder != order {
			t.Errorf(
				"FAIL: Want correct ordering for route %v " +
				"picking up %v to be: %v but we got %v",
				input.driverRoute,
				input.riderRoute,
				input.expectedOrder,
				order,
			)
		}
	}
}

var haversineDistanceTestCases = []struct {
	coordA         coord
	coordB         coord
	expectedMeters float64
}{
	{
		coord{lat: 38.89768, lng: -77.03653},
		coord{lat: 38.89736, lng: -77.04173},
		451,
	},
	{
		coord{lat: 51.510357, lng: -0.116773},
		coord{lat: 38.889931, lng: -77.009003},
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
		driverID, matched := weightedInverseRandRider(
			input.pickupSet,
			input.prevMatches,
			input.randSeed,
		)

		if input.expectedID != driverID || input.expectedOK != matched {
			t.Errorf(
				"FAIL: Want match to be: (%v, %v) but we got " +
				"(%v, %v)",
				input.expectedID,
				input.expectedOK,
				driverID,
				matched,
			)
		}
	}
}
