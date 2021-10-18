package _993_cousins_in_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsCousins(t *testing.T) {
	testCases := []struct {
		got  []int
		gotX int
		gotY int
		want bool
	}{
		{
			got:  []int{1, 2, 3, 4},
			gotX: 4,
			gotY: 3,
			want: false,
		},
		{
			got:  []int{1, 2, 3, NULL, 4, NULL, 5},
			gotX: 5,
			gotY: 4,
			want: true,
		},
		{
			got:  []int{1, 2, 3, NULL, 4},
			gotX: 2,
			gotY: 3,
			want: false,
		},
		{
			got:  []int{1, 2, 3, NULL, NULL, NULL, 4, 5},
			gotX: 1,
			gotY: 2,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isCousins(BinaryTreeFromSlice(testCase.got), testCase.gotX, testCase.gotY)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
