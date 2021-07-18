package _025_reverse_nodes_in_k_group

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseKGroup(t *testing.T) {
	testCases := []struct {
		gotHead []int
		gotK    int
		want    []int
	}{
		{
			gotHead: []int{1, 2, 3, 4, 5},
			gotK:    2,
			want:    []int{2, 1, 4, 3, 5},
		},
		{
			gotHead: []int{1, 2, 3, 4, 5},
			gotK:    3,
			want:    []int{3, 2, 1, 4, 5},
		},
		{
			gotHead: []int{1, 2, 3, 4, 5},
			gotK:    1,
			want:    []int{1, 2, 3, 4, 5},
		},
		{
			gotHead: []int{1},
			gotK:    1,
			want:    []int{1},
		},
	}

	for _, testCase := range testCases {
		actual := reverseKGroup(LinkedListFromSlice(testCase.gotHead), testCase.gotK)

		assert.Check(t, is.DeepEqual(actual.ToSlice(), testCase.want))
	}
}
