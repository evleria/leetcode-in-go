package _732_find_the_highest_altitude

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLargestAltitude(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{-5, 1, 5, 0, -7},
			want:    1,
		},
		{
			gotNums: []int{-4, -3, -2, -1, 4, 3, 2},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := largestAltitude(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
