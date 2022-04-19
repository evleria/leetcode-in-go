package _099_recover_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRecoverTree(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 3, NULL, NULL, 2},
			want: []int{3, 1, NULL, NULL, 2},
		},
		{
			got:  []int{3, 1, 4, NULL, NULL, 2},
			want: []int{2, 1, 4, NULL, NULL, 3},
		},
	}

	for _, testCase := range testCases {
		tree := BinaryTreeFromSlice(testCase.got)
		recoverTree(tree)
		assert.Assert(t, is.DeepEqual(tree.ToSlice(), testCase.want))
	}
}
