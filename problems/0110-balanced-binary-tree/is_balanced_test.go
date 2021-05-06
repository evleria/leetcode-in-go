package _110_balanced_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{3, 9, 20, NULL, NULL, 15, 7},
			want: true,
		},
		{
			got:  []int{1, 2, 2, 3, 3, NULL, NULL, 4, 4},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isBalanced(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
