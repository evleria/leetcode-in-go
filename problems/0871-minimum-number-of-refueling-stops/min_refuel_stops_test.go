package _871_minimum_number_of_refueling_stops

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinRefuelStops(t *testing.T) {
	testCases := []struct {
		gotTarget    int
		gotStartFuel int
		gotStations  [][]int
		want         int
	}{
		{
			gotTarget:    1,
			gotStartFuel: 1,
			gotStations:  [][]int{},
			want:         0,
		},
		{
			gotTarget:    100,
			gotStartFuel: 1,
			gotStations:  [][]int{{10, 100}},
			want:         -1,
		},
		{
			gotTarget:    100,
			gotStartFuel: 10,
			gotStations: [][]int{
				{10, 60},
				{20, 30},
				{30, 30},
				{60, 40}},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := minRefuelStops(testCase.gotTarget, testCase.gotStartFuel, testCase.gotStations)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
