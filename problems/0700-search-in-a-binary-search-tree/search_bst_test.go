package _700_search_in_a_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSearchBST(t *testing.T) {
	testCases := []struct {
		got    []int
		gotVal int
		want   []int
	}{
		{
			got:    []int{4, 2, 7, 1, 3},
			gotVal: 2,
			want:   []int{2, 1, 3},
		},
		{
			got:    []int{4, 2, 7, 1, 3},
			gotVal: 5,
			want:   []int{},
		},
	}

	for _, testCase := range testCases {
		actual := searchBST(BinaryTreeFromSlice(testCase.got), testCase.gotVal).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
