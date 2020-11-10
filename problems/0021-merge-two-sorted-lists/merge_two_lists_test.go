package _021_merge_two_sorted_lists

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		gotList1 []int
		gotList2 []int
		want     []int
	}{
		{
			gotList1: []int{1, 2},
			gotList2: []int{1, 3, 4},
			want:     []int{1, 1, 2, 3, 4},
		},
	}

	for _, testCase := range testCases {
		actual := mergeTwoLists(FromSlice(testCase.gotList1), FromSlice(testCase.gotList2))

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want))
	}
}
