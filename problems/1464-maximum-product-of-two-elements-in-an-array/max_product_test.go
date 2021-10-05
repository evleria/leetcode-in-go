package _464_maximum_product_of_two_elements_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 4, 5, 2},
			want: 12,
		},
		{
			got:  []int{1, 5, 4, 5},
			want: 16,
		},
		{
			got:  []int{3, 7},
			want: 12,
		},
	}

	for _, testCase := range testCases {
		actual := maxProduct(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
