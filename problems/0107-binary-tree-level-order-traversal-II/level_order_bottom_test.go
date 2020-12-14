package _107_binary_tree_level_order_traversal_II

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLevelOrderBottom(t *testing.T) {
	testCases := []struct {
		got  []int
		want [][]int
	}{
		{
			got: []int{3, 9, 20, NULL, NULL, 15, 7},
			want: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
	}

	for _, testCase := range testCases {
		actual := levelOrderBottom(FromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
