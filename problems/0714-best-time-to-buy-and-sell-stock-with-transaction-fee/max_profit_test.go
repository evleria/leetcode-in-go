package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		got    []int
		gotFee int
		want   int
	}{
		{
			got:    []int{1, 3, 2, 8, 4, 9},
			gotFee: 2,
			want:   8,
		},
		{
			got:    []int{1, 3, 7, 5, 10, 3},
			gotFee: 3,
			want:   6,
		},
		{
			got:    []int{9, 8, 7, 1, 2},
			gotFee: 3,
			want:   0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProfit(testCase.got, testCase.gotFee)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
