package _129_sum_root_to_leaf_numbers

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumNumbers(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 2, 3},
			want: 25,
		},
		{
			got:  []int{4, 9, 0, 5, 1},
			want: 1026,
		},
	}

	for _, testCase := range testCases {
		actual := sumNumbers(FromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
