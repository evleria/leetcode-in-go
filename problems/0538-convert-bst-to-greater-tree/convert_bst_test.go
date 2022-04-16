package _538_convert_bst_to_greater_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConvertBST(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{4, 1, 6, 0, 2, 5, 7, NULL, NULL, NULL, 3, NULL, NULL, NULL, 8},
			want: []int{30, 36, 21, 36, 35, 26, 15, NULL, NULL, NULL, 33, NULL, NULL, NULL, 8},
		},
		{
			got:  []int{0, NULL, 1},
			want: []int{1, NULL, 1},
		},
	}

	for _, testCase := range testCases {
		actual := convertBST(BinaryTreeFromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
