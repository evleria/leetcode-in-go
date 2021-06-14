package _710_maximum_units_on_a_truck

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximumUnits(t *testing.T) {
	testCases := []struct {
		gotBox   [][]int
		gotTruck int
		want     int
	}{
		{
			gotBox:   [][]int{{1, 3}, {2, 2}, {3, 1}},
			gotTruck: 4,
			want:     8,
		},
		{
			gotBox:   [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
			gotTruck: 10,
			want:     91,
		},
	}

	for _, testCase := range testCases {
		actual := maximumUnits(testCase.gotBox, testCase.gotTruck)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
