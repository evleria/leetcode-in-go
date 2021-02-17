package _145_binary_tree_postorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPostorderTraversal(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, NULL, 2, 3},
			want: []int{3, 2, 1},
		},
	}

	for _, testCase := range testCases {
		actual := postorderTraversal(FromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
