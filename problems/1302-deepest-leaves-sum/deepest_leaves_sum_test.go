package _302_deepest_leaves_sum

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDeepLeavesSum(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 2, 3, 4, 5, NULL, 6, 7, NULL, NULL, NULL, NULL, 8},
			want: 15,
		},
		{
			got:  []int{6, 7, 8, 2, 7, 1, 3, 9, NULL, 1, 4, NULL, NULL, NULL, 5},
			want: 19,
		},
	}

	for _, testCase := range testCases {
		actual := deepestLeavesSum(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
