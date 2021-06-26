package _643_maximum_average_subarray_I

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMaxAverage(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      float64
	}{
		{
			gotNums:   []int{1, 12, -5, -6, 50, 3},
			gotTarget: 4,
			want:      12.75000,
		},
		{
			gotNums:   []int{5},
			gotTarget: 1,
			want:      5.00000,
		},
	}

	for _, testCase := range testCases {
		actual := findMaxAverage(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
