package _222_count_complete_tree_nodes

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountNodes(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 2, 3, 4, 5, 6},
			want: 6,
		},
		{
			got:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: 8,
		},
		{
			got:  []int{},
			want: 0,
		},
		{
			got:  []int{1},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := countNodes(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
