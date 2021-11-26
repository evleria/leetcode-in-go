package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		gotI []int
		gotP []int
		want []int
	}{
		{
			gotI: []int{9, 3, 15, 20, 7},
			gotP: []int{9, 15, 7, 20, 3},
			want: []int{3, 9, 20, NULL, NULL, 15, 7},
		},
		{
			gotI: []int{-1},
			gotP: []int{-1},
			want: []int{-1},
		},
	}

	for _, testCase := range testCases {
		actual := buildTree(testCase.gotI, testCase.gotP)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want))
	}
}
