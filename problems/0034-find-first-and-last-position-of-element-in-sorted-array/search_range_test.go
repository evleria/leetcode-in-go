package _034_find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      []int
	}{
		{
			gotNums:   []int{5, 7, 7, 8, 8, 10},
			gotTarget: 8,
			want:      []int{3, 4},
		},
		{
			gotNums:   []int{5, 7, 7, 8, 8, 10},
			gotTarget: 6,
			want:      []int{-1, -1},
		},
		{
			gotNums:   []int{},
			gotTarget: 0,
			want:      []int{-1, -1},
		},
	}

	for _, testCase := range testCases {
		actual := searchRange(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
