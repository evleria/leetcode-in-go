package _814_binary_tree_pruning

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPruneTree(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, NULL, 0, 0, 1},
			want: []int{1, NULL, 0, NULL, 1},
		},
		{
			got:  []int{1, 0, 1, 0, 0, 0, 1},
			want: []int{1, NULL, 1, NULL, 1},
		},
	}

	for _, testCase := range testCases {
		actual := pruneTree(BinaryTreeFromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
