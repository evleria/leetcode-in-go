package _563_binary_tree_tilt

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindTilt(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 2, 3},
			want: 1,
		},
		{
			got:  []int{4, 2, 9, 3, 5, NULL, 7},
			want: 15,
		},
		{
			got:  []int{21, 7, 14, 1, 1, 2, 2, 3, 3},
			want: 9,
		},
	}

	for _, testCase := range testCases {
		actual := findTilt(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
