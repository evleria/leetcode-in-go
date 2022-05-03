package _581_shortest_unsorted_continuous_subarray

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBackspaceCompare(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{2, 6, 4, 8, 10, 9, 15},
			want: 5,
		},
		{
			got:  []int{1, 2, 3, 4},
			want: 0,
		},
		{
			got:  []int{1},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := findUnsortedSubarray(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
