package _598_range_addition_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxCount(t *testing.T) {
	testCases := []struct {
		gotM   int
		gotN   int
		gotOps [][]int
		want   int
	}{
		{
			gotM:   3,
			gotN:   3,
			gotOps: [][]int{{2, 2}, {3, 3}},
			want:   4,
		},
		{
			gotM:   3,
			gotN:   3,
			gotOps: [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}},
			want:   4,
		},
		{
			gotM:   3,
			gotN:   3,
			gotOps: [][]int{},
			want:   9,
		},
	}

	for _, testCase := range testCases {
		actual := maxCount(testCase.gotM, testCase.gotN, testCase.gotOps)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotM)
	}
}
