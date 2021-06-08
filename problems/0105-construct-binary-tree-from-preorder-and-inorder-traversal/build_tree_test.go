package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"fmt"
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		gotPreorder []int
		gotInorder  []int
		want        []int
	}{
		{
			gotPreorder: []int{3, 9, 20, 15, 7},
			gotInorder:  []int{9, 3, 15, 20, 7},
			want:        []int{3, 9, 20, NULL, NULL, 15, 7},
		},
		{
			gotPreorder: []int{-1},
			gotInorder:  []int{-1},
			want:        []int{-1},
		},
	}

	for _, testCase := range testCases {
		actual := buildTree(testCase.gotPreorder, testCase.gotInorder).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%v", testCase))
	}
}
