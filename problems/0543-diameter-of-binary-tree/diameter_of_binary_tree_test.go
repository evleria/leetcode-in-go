package _543_diameter_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 2, 3, 4, 5},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := diameterOfBinaryTree(FromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}