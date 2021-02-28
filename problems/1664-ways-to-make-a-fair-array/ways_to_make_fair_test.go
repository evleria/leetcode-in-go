package _664_ways_to_make_a_fair_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWaysToMakeFair(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{2, 1, 6, 4},
			want:    1,
		},
		{
			gotNums: []int{1, 1, 1},
			want:    3,
		},
		{
			gotNums: []int{4, 1, 1, 2, 5, 1, 5, 4},
			want:    2,
		},
	}

	for _, testCase := range testCases {
		actual := waysToMakeFair(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
