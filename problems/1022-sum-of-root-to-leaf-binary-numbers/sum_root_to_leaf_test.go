package _022_sum_of_root_to_leaf_binary_numbers

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumRootToLeaf(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 0, 1, 0, 1, 0, 1},
			want: 22,
		},
		{
			got:  []int{0},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := sumRootToLeaf(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
