package _662_maximum_width_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWidthOfBinaryTree(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 3, 2, 5, 3, NULL, 9},
			want: 4,
		},
		{
			got:  []int{1, 3, NULL, 5, 3},
			want: 2,
		},
		{
			got:  []int{1, 3, 2, 5},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := widthOfBinaryTree(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
