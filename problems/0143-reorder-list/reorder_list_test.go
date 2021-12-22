package _143_reorder_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReorderList(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 3, 4},
			want: []int{1, 4, 2, 3},
		},
		{
			got:  []int{1, 2, 3, 4, 5},
			want: []int{1, 5, 2, 4, 3},
		},
	}

	for _, testCase := range testCases {
		list := LinkedListFromSlice(testCase.got)
		reorderList(list)

		assert.Check(t, is.DeepEqual(list.ToSlice(), testCase.want), testCase.got)
	}
}
