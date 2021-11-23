package _952_largest_component_size_by_common_factor

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLargestComponentSize(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{4, 6, 15, 35},
			want:    4,
		},
		{
			gotNums: []int{20, 50, 9, 63},
			want:    2,
		},
		{
			gotNums: []int{2, 3, 6, 7, 4, 12, 21, 39},
			want:    8,
		},
	}

	for _, testCase := range testCases {
		actual := largestComponentSize(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
