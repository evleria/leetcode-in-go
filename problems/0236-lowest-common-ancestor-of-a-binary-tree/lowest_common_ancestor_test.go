package _236_lowest_common_ancestor_of_a_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLowestCommonAncestor(t *testing.T) {
	testCases := []struct {
		gotTree []int
		gotP    int
		gotQ    int
		want    int
	}{
		{
			gotTree: []int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4},
			gotP:    5,
			gotQ:    1,
			want:    3,
		},
		{
			gotTree: []int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4},
			gotP:    5,
			gotQ:    4,
			want:    5,
		},
	}

	for _, testCase := range testCases {
		tree := BinaryTreeFromSlice(testCase.gotTree)
		actual := lowestCommonAncestor(tree, tree.FindInBT(testCase.gotP), tree.FindInBT(testCase.gotQ))

		assert.Check(t, is.Equal(actual.Val, testCase.want))
	}
}
