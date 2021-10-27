package _257_binary_tree_paths

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBinaryTreePaths(t *testing.T) {
	testCases := []struct {
		got  []int
		want []string
	}{
		{
			got:  []int{1, 2, 3, NULL, 5},
			want: []string{"1->2->5", "1->3"},
		},
		{
			got:  []int{1},
			want: []string{"1"},
		},
	}

	for _, testCase := range testCases {
		actual := binaryTreePaths(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
