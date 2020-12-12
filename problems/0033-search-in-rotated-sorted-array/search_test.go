package _033_search_in_rotated_sorted_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{4, 5, 6, 7, 0, 1, 2},
			gotTarget: 0,
			want:      4,
		},
		{
			gotNums:   []int{4, 5, 6, 7, 0, 1, 2},
			gotTarget: 3,
			want:      -1,
		},
		{
			gotNums:   []int{1},
			gotTarget: 0,
			want:      -1,
		},
	}

	for _, testCase := range testCases {
		actual := search(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
