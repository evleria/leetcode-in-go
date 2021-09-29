package _725_split_linked_list_in_parts

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSplitListToParts(t *testing.T) {
	testCases := []struct {
		gotList []int
		gotK    int
		want    [][]int
	}{
		{
			gotList: []int{1, 2, 3},
			gotK:    5,
			want:    [][]int{{1}, {2}, {3}, {}, {}},
		},
		{
			gotList: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			gotK:    3,
			want:    [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
		},
	}

	for _, testCase := range testCases {
		actual := splitListToParts(LinkedListFromSlice(testCase.gotList), testCase.gotK)
		slices := make([][]int, len(actual))
		for i, list := range actual {
			slices[i] = list.ToSlice()
		}

		assert.Check(t, is.DeepEqual(slices, testCase.want))
	}
}
