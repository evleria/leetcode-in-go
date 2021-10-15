package _309_best_time_to_buy_and_sell_stock_with_cooldown

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
			got:  []int{1, 2, 3, 0, 2},
			want: 3,
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
