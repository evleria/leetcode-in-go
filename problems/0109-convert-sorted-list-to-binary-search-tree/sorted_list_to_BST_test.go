package _109_convert_sorted_list_to_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortedListToBST(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{-10, -3, 0, 5, 9},
			want: []int{0, -3, 9, -10, NULL, 5},
		},
		{
			got:  []int{},
			want: []int{},
		},
		{
			got:  []int{1, 3},
			want: []int{3, 1},
		},
	}

	for _, testCase := range testCases {
		actual := sortedListToBST(LinkedListFromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
