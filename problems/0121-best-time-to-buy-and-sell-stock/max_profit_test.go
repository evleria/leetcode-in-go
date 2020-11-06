package _121_best_time_to_buy_and_sell_stock

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
			got:  []int{},
			want: 0,
		},
		{
			got:  []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			got:  []int{7, 6, 4, 3, 1},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProfit(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
