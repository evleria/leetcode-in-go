package _897_increasing_order_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIncreasingBST(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{5, 1, 7},
			want: []int{1, NULL, 5, NULL, 7},
		},
		{
			got:  []int{5, 3, 6, 2, 4, NULL, 8, 1, NULL, NULL, NULL, 7, 9},
			want: []int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5, NULL, 6, NULL, 7, NULL, 8, NULL, 9},
		},
	}

	for _, testCase := range testCases {
		actual := increasingBST(BinaryTreeFromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
