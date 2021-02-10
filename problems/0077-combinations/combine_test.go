package _077_combinations

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCombine(t *testing.T) {
	testCases := []struct {
		gotN int
		gotK int
		want [][]int
	}{
		{
			gotN: 4,
			gotK: 2,
			want: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			gotN: 1,
			gotK: 1,
			want: [][]int{
				{1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := combine(testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
