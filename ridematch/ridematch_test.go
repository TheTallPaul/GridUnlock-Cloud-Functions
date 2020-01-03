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
				`FAIL: Want distance from %v to %v to be: %v but
				we got %v`,
				input.coordA,
				input.coordB,
				input.expectedMeters,
				meters,
			)
		}
	}
}

var weightedInverseRandRiderTestCases = []struct {
	pickupSet possiblePickups

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
					distance: 6.0,
				},
			},
		},
	},
	{
		possiblePickups{
			riderID: "rider1",
			networkRiderDistances: []riderDistance {
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
		},
	},
}


func TestWeightedInverseRandRider(t *testing.T) {

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
				`FAIL: Want %v inverted to be length %v but we
				got %v`,
				input.slice,
				len(input.expectedSlice),
				invertedSlice,
			)
		} else {
			for i, x := range invertedSlice {
				if x != input.expectedSlice[i] {
					t.Errorf(
						`FAIL: Want %v inverted to be %v
						but we got %v`,
						input.slice,
						input.expectedSlice,
						invertedSlice,
					)
				}
			}
		}
	}
}
