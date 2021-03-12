package _322_coin_change

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCoinChange(t *testing.T) {
	testCases := []struct {
		got       []int
		gotAmount int
		want      int
	}{
		{
			got:       []int{1, 2},
			gotAmount: 2,
			want:      1,
		},
		{
			got:       []int{1, 2, 5},
			gotAmount: 11,
			want:      3,
		},
		{
			got:       []int{2},
			gotAmount: 3,
			want:      -1,
		},
		{
			got:       []int{1},
			gotAmount: 0,
			want:      0,
		},
		{
			got:       []int{1},
			gotAmount: 2,
			want:      2,
		},
	}

	for _, testCase := range testCases {
		actual := coinChange(testCase.got, testCase.gotAmount)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
