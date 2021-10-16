package _123_best_time_to_buy_and_sell_stock_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 3, 5, 0, 0, 3, 1, 4},
			want: 6,
		},
		{
			got:  []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			got:  []int{7, 6, 4, 3, 1},
			want: 0,
		},
		{
			got:  []int{1},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProfit(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
