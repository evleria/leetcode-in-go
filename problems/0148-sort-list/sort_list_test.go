package _148_sort_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortList(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{4, 2, 1, 3},
			want: []int{1, 2, 3, 4},
		},
		{
			got:  []int{-1, 5, 3, 4, 0},
			want: []int{-1, 0, 3, 4, 5},
		},
		{
			got:  []int{},
			want: []int{},
		},
	}

	for _, testCase := range testCases {
		list := LinkedListFromSlice(testCase.got)
		actual := sortList(list)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want), testCase.got)
	}
}
