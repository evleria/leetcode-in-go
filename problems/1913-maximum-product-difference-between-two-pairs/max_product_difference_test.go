package _913_maximum_product_difference_between_two_pairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProductDifference(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{5, 6, 2, 7, 4},
			want:    34,
		},
		{
			gotNums: []int{4, 2, 5, 9, 7, 4, 8},
			want:    64,
		},
	}

	for _, testCase := range testCases {
		actual := maxProductDifference(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
