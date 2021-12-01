package _057_smallest_index_with_equal_value

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSmallestEqual(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{0, 1, 2},
			want: 0,
		},
		{
			got:  []int{4, 3, 2, 1},
			want: 2,
		},
		{
			got:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			want: -1,
		},
		{
			got:  []int{2, 1, 3, 5, 2},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := smallestEqual(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
