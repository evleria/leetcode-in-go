package _475_final_prices_with_a_special_discount_in_a_shop

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFinalPrices(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{8, 4, 6, 2, 3},
			want: []int{4, 2, 4, 2, 3},
		},
		{
			got:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			got:  []int{10, 1, 1, 6},
			want: []int{9, 0, 1, 6},
		},
	}

	for _, testCase := range testCases {
		actual := finalPrices(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
