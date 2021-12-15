package _147_insertion_sort_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestInsertionSortList(t *testing.T) {
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
	}

	for _, testCase := range testCases {
		actual := insertionSortList(LinkedListFromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
