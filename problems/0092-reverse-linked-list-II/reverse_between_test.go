package _092_reverse_linked_list_II

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseBetween(t *testing.T) {
	testCases := []struct {
		got      []int
		gotLeft  int
		gotRight int
		want     []int
	}{
		{
			got:      []int{1, 2, 3, 4, 5},
			gotLeft:  2,
			gotRight: 4,
			want:     []int{1, 4, 3, 2, 5},
		},
		{
			got:      []int{1, 2, 3},
			gotLeft:  1,
			gotRight: 2,
			want:     []int{2, 1, 3},
		},
		{
			got:      []int{5},
			gotLeft:  1,
			gotRight: 1,
			want:     []int{5},
		},
	}

	for _, testCase := range testCases {
		actual := reverseBetween(LinkedListFromSlice(testCase.got), testCase.gotLeft, testCase.gotRight)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want), testCase.got)
	}
}
