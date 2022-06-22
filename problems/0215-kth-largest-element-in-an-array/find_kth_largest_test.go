package _215_kth_largest_element_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotX    int
		want    int
	}{
		{
			gotNums: []int{3, 2, 1, 5, 6, 4},
			gotX:    2,
			want:    5,
		},
		{
			gotNums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			gotX:    4,
			want:    4,
		},
	}

	for _, testCase := range testCases {
		actual := findKthLargest(testCase.gotNums, testCase.gotX)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
