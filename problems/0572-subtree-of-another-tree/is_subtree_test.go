package _572_subtree_of_another_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSubtree(t *testing.T) {
	testCases := []struct {
		gotRoot    []int
		gotSubtree []int
		want       bool
	}{
		{
			gotRoot:    []int{3, 4, 5, 1, 2},
			gotSubtree: []int{4, 1, 2},
			want:       true,
		},
		{
			gotRoot:    []int{3, 4, 5, 1, 2, NULL, NULL, 0},
			gotSubtree: []int{4, 1, 2},
			want:       false,
		},
	}

	for _, testCase := range testCases {
		actual := isSubtree(BinaryTreeFromSlice(testCase.gotRoot), BinaryTreeFromSlice(testCase.gotSubtree))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotRoot, testCase.gotSubtree)
	}
}
