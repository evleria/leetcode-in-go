package _540_single_element_in_a_sorted_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSingleNonDuplicate(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			want:    2,
		},
		{
			gotNums: []int{3, 3, 7, 7, 10, 11, 11},
			want:    10,
		},
	}

	for _, testCase := range testCases {
		actual := singleNonDuplicate(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
