package _671_second_minimum_node_in_a_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindSecondMinimumValue(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{2, 2, 5, NULL, NULL, 5, 7},
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := findSecondMinimumValue(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
