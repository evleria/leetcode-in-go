package _128_longest_consecutive_sequence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestConsecutive(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{100, 4, 200, 1, 3, 2},
			want:    4,
		},
		{
			gotNums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want:    9,
		},
	}

	for _, testCase := range testCases {
		actual := longestConsecutive(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
