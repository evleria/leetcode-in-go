package _304_find_n_unique_integers_sum_up_to_zero

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumZero(t *testing.T) {
	testCases := []struct {
		got  int
		want []int
	}{
		{
			got:  5,
			want: []int{-10, 1, 2, 3, 4},
		},
		{
			got:  3,
			want: []int{-3, 1, 2},
		},
		{
			got:  1,
			want: []int{0},
		},
	}

	for _, testCase := range testCases {
		actual := sumZero(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
