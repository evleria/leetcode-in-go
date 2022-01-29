package _084_largest_rectangle_in_histogram

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLargestRectangleInHistogram(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{2, 1, 5, 6, 2, 3},
			want: 10,
		},
		{
			got:  []int{2, 4},
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := largestRectangleArea(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
