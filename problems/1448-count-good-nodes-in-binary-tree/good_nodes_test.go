package _448_count_good_nodes_in_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGoodNodes(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 1, 4, 3, NULL, 1, 5},
			want: 4,
		},
		{
			got:  []int{3, 3, NULL, 4, 2},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := goodNodes(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
