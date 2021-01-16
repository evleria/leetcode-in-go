package _152_maximum_product_subarray

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{2, 3, -2, 4},
			want:    6,
		},
		{
			gotNums: []int{-2, 0, -1},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProduct(testCase.gotNums)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
