package _083_remove_duplicates_from_sorted_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			got:  []int{1, 1, 2},
			want: []int{1, 2},
		},
	}

	for _, testCase := range testCases {
		actual := deleteDuplicates(FromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
