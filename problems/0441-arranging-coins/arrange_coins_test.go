package _441_arranging_coins

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestArrangeCoins(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  5,
			want: 2,
		},
		{
			got:  8,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := arrangeCoins(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
