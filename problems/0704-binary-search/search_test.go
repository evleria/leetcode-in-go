package _704_binary_search

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
			gotNums:   []int{-1, 0, 3, 5, 9, 12},
			gotTarget: 9,
			want:      4,
		},
		{
			gotNums:   []int{-1, 0, 3, 5, 9, 12},
			gotTarget: 2,
			want:      -1,
		},
		{
			gotNums:   []int{5},
			gotTarget: 5,
			want:      0,
		},
	}

	for _, testCase := range testCases {
		actual := search(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
