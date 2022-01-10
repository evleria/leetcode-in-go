package _312_burst_balloons

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxCoins(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 1, 5, 8},
			want: 167,
		},
		{
			got:  []int{1, 5},
			want: 10,
		},
	}

	for _, testCase := range testCases {
		actual := maxCoins(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
