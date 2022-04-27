package _202_smallest_string_with_swaps

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSmallestStringWithSwaps(t *testing.T) {
	testCases := []struct {
		gotS     string
		gotPairs [][]int
		want     string
	}{
		{
			gotS:     "dcab",
			gotPairs: [][]int{{0, 3}, {1, 2}},
			want:     "bacd",
		},
		{
			gotS:     "dcab",
			gotPairs: [][]int{{0, 3}, {1, 2}, {0, 2}},
			want:     "abcd",
		},
		{
			gotS:     "cba",
			gotPairs: [][]int{{0, 1}, {1, 2}},
			want:     "abc",
		},
	}

	for _, testCase := range testCases {
		actual := smallestStringWithSwaps(testCase.gotS, testCase.gotPairs)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
