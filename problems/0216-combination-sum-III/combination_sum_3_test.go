package _216_combination_sum_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCombinationSum3(t *testing.T) {
	testCases := []struct {
		gotK int
		gotN int
		want [][]int
	}{
		{
			gotK: 3,
			gotN: 7,
			want: [][]int{{1, 2, 4}},
		},
		{
			gotK: 3,
			gotN: 9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			gotK: 4,
			gotN: 1,
			want: [][]int{},
		},
	}

	for _, testCase := range testCases {
		actual := combinationSum3(testCase.gotK, testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
