package _672_richest_customer_wealth

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximumWealth(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{
				{1, 2, 3},
				{3, 2, 1},
			},
			want: 6,
		},
		{
			gotNums: [][]int{
				{1, 5},
				{7, 3},
				{3, 5},
			},
			want: 10,
		},
	}

	for _, testCase := range testCases {
		actual := maximumWealth(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
