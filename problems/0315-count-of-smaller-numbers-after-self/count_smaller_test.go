package _315_count_of_smaller_numbers_after_self

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountSmaller(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{5, 2, 6, 1},
			want:    []int{2, 1, 1, 0},
		},
		{
			gotNums: []int{-1},
			want:    []int{0},
		},
		{
			gotNums: []int{-1, -1},
			want:    []int{0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := countSmaller(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
