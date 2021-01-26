package _235_lowest_common_ancestor_of_a_binary_search_tree

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
			gotTree: []int{6, 2, 8, 0, 4, 7, 9, NULL, NULL, 3, 5},
			gotP:    2,
			gotQ:    8,
			want:    6,
		},
		{
			gotTree: []int{6, 2, 8, 0, 4, 7, 9, NULL, NULL, 3, 5},
			gotP:    2,
			gotQ:    4,
			want:    2,
		},
	}

	for _, testCase := range testCases {
		tree := FromSlice(testCase.gotTree)
		actual := lowestCommonAncestor(tree, tree.FindInBST(testCase.gotP), tree.FindInBST(testCase.gotQ))

		assert.Check(t, is.Equal(actual.Val, testCase.want))
	}
}
