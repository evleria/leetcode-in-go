package _532_k_diff_pairs_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindPairs(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    int
	}{
		{
			gotNums: []int{1, 3, 1, 5, 4},
			gotK:    0,
			want:    1,
		},
		{
			gotNums: []int{3, 1, 4, 1, 5},
			gotK:    2,
			want:    2,
		},
		{
			gotNums: []int{1, 2, 3, 4, 5},
			gotK:    1,
			want:    4,
		},
	}

	for _, testCase := range testCases {
		actual := findPairs(testCase.gotNums, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
