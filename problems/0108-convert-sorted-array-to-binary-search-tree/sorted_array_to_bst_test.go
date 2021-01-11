package _108_convert_sorted_array_to_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestInorderTraversal(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{-10, -3, 0, 5, 9},
			want: []int{0, -3, 9, -10, NULL, 5},
		},
	}

	for _, testCase := range testCases {
		actual := sortedArrayToBST(testCase.got).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
