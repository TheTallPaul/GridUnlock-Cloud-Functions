package ridematch_test

import (
    "testing"
)

var haversineDistanceTestCases = []struct {
	coordA         coord
	coordB         coord
	expectedMeters float64
}{
	{
		coord{lat: 38.89768, lon: -77.03653},
		coord{lat: 38.89736, lon: -77.04173},
		451.41,
	},
	{
		coord{lat: 51.510357, lon: -0.116773},
		coord{lat: 38.889931, lon: -77.009003},
		5897658.288856054,
	}
}

func TestHaversineDistance(t *testing.T) {
	for _, input := range haversineDistanceTestCases {
		meters := haversineDistance(input.coordA, input.coordB)

		if input.expectedMeters != meters {
			t.Errorf(
				"FAIL: Want distance from %v to %v to be: %v but
				we got %v",
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
		pickupSet{
			riderID: "rider1",
			networkRiderDistances: [
				riderDistance{
					driverID: "driver1",
					distance: 1.0

				},
				riderDistance{
					driverID: "driver2",
					distance: 6.0
				}
			]
		}
	}
}


func TestWeightedInverseRandRider(t *testing.T) {

}

var inverseArrayTestCases = []struct {
	array []float64
	expectedArray []float64
}{
	{
		[1.0, 0.5, 22.3],
		[1.0, 2.0, 0.044843],
	},
	{
		[-0.2, -1.5, 4.0],
		[-5.0, 0.6666666666666666, 0.25],
	}
}

func TestInverseArray(t *testing.T) {
	for _, input := range inverseArrayTestCases {
		invertedArray := inverseArray(input.array)

		if input.expectedArray != invertedArray {
			t.Errorf(
				"FAIL: Want %v inverted to be %v but we got %v",
				input.array,
				input.expectedArray,
				input.invertedArray,
			)
		}

	}
}
