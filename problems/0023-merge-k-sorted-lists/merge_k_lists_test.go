package _023_merge_k_sorted_lists

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMergeKLists(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want []int
	}{
		{
			got:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			got:  [][]int{{}},
			want: []int{},
		},
		{
			got:  [][]int{{}, {}},
			want: []int{},
		},
	}

	for _, testCase := range testCases {
		lists := make([]*ListNode, 0, len(testCase.got))
		for _, l := range testCase.got {
			lists = append(lists, LinkedListFromSlice(l))
		}
		actual := mergeKLists(lists)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want))
	}
}
