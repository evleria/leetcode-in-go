package _501_find_mode_in_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMode(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, NULL, 2, 2},
			want: []int{2},
		},
		{
			got:  []int{1, NULL, 2},
			want: []int{1, 2},
		},
		{
			got:  []int{0},
			want: []int{0},
		},
	}

	for _, testCase := range testCases {
		actual := findMode(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
