package _008_construct_binary_search_tree_from_preorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBstFromPreorder(t *testing.T) {
	testCases := []struct {
		gotTree []int
		want    []int
	}{
		{
			gotTree: []int{8, 5, 1, 7, 10, 12},
			want:    []int{8, 5, 10, 1, 7, NULL, 12},
		},
		{
			gotTree: []int{1, 3},
			want:    []int{1, NULL, 3},
		},
	}

	for _, testCase := range testCases {
		actual := bstFromPreorder(testCase.gotTree)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want))
	}
}
