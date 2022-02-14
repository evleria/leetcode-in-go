package _852_peak_index_in_a_mountain_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{0, 1, 0},
			want:    1,
		},
		{
			gotNums: []int{0, 2, 1, 0},
			want:    1,
		},
		{
			gotNums: []int{0, 10, 5, 2},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := peakIndexInMountainArray(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
