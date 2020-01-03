package ridematch

import (
	"testing"
)

var haversineDistanceTestCases = []struct {
	coordA         coord
	coordB         coord
	expectedMeters float64
}{
	{
		coord{lat: 38.89768, lng: -77.03653},
		coord{lat: 38.89736, lng: -77.04173},
		451.4110744692723,
	},
	{
		coord{lat: 51.510357, lng: -0.116773},
		coord{lat: 38.889931, lng: -77.009003},
		5897658.288856054,
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
	pickupSet     possiblePickups
	randSeed      int64
	expectedMatch riderDriverMatch
}{
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: []riderDistance {
				riderDistance{
					driverID: "driver1",
					distance: 1.0,
				},
				riderDistance{
					driverID: "driver2",
					distance: 100.0,
				},
			},
		},
		1,
		riderDriverMatch{
			riderID: "rider1",
			driverID: "driver1",
		},
	},
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: []riderDistance {
				riderDistance{
					driverID: "driver1",
					distance: 430.1,
				},
				riderDistance{
					driverID: "driver2",
					distance: 3.0,
				},
				riderDistance{
					driverID: "driver3",
					distance: 43.0,
				},
			},
		},
		1,
		riderDriverMatch{
			riderID: "rider1",
			driverID: "driver2",
		},
	},
}


func TestWeightedInverseRandRider(t *testing.T) {
	for _, input := range weightedInverseRandRiderTestCases {
		match := weightedInverseRandRider(
			input.pickupSet,
			input.randSeed,
		)

		if input.expectedMatch.driverID != match.driverID ||
			input.expectedMatch.riderID  != match.riderID {
			t.Errorf(
				"FAIL: Want match to be: %v but we got %v",
				input.expectedMatch,
				match,
			)
		}
	}
}

var convertPickupsToSliceTestCases = []struct {
	distances     []riderDistance
	expectedSlice []float64
}{
	{
		[]riderDistance{
			riderDistance{
				driverID: "driver1",
				distance: 1.0,
			},
			riderDistance{
				driverID: "driver2",
				distance: 6.0,
			},
		},
		[]float64{1.0, 6.0,},

	},
	{
		[]riderDistance{
			riderDistance{
				driverID: "driver1",
				distance: 0.1,
			},
			riderDistance{
				driverID: "driver2",
				distance: 2.0,
			},
			riderDistance{
				driverID: "driver3",
				distance: 43.0,
			},
		},
		[]float64{0.1, 2.0, 43.0,},
	},
}

func TestConvertPickupsToSlice(t *testing.T) {
	for _, input := range convertPickupsToSliceTestCases {
		slice := convertPickupsToSlice(input.distances)

		if  len(slice) != len(input.expectedSlice) {
			t.Errorf(
				"FAIL: Want slice to be length: %v but we " +
				"got %v",
				len(input.expectedSlice),
				slice,
			)
		} else {
			for i, x := range slice {
				if x != input.expectedSlice[i] {
					t.Errorf(
						"FAIL: Want %v inverted to " +
						"be: %v but we got %v",
						input.distances,
						input.expectedSlice,
						slice,
					)
				}
			}
		}
	}
}

var sumSliceTestCases = []struct {
	slice         []float64
	expectedTotal float64
}{
	{
		[]float64{0.0, 2.0, -0.0},
		2.0,
	},
	{
		[]float64{1, 2, -3},
		0.0,
	},
}

func TestSumSlice(t *testing.T) {
	for _, input := range sumSliceTestCases {
		total := sumSlice(input.slice)

		if input.expectedTotal != total {
			t.Errorf(
				"FAIL: Want sum from %v to be: %v but we got " +
				"%v",
				input.slice,
				input.expectedTotal,
				total,
			)
		}
	}
}

var inverseSliceTestCases = []struct {
	slice         []float64
	expectedSlice []float64
}{
	{
		[]float64{1.0, 0.5, 22.3},
		[]float64{1.0, 2.0, 0.04484304932735426},
	},
	{
		[]float64{-0.2, -1.5, 4.0},
		[]float64{-5.0, -0.6666666666666666, 0.25},
	},
}

func TestInverseSlice(t *testing.T) {
	for _, input := range inverseSliceTestCases {
		invertedSlice := inverseSlice(input.slice)

		if  len(invertedSlice) != len(input.expectedSlice) {
			t.Errorf(
				"FAIL: Want %v inverted to be length: %v but " +
				"we got %v",
				input.slice,
				len(input.expectedSlice),
				invertedSlice,
			)
		} else {
			for i, x := range invertedSlice {
				if x != input.expectedSlice[i] {
					t.Errorf(
						"FAIL: Want %v inverted to " +
						"be: %v but we got %v",
						input.slice,
						input.expectedSlice,
						invertedSlice,
					)
				}
			}
		}
	}
}
