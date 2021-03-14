package _721_swapping_nodes_in_a_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSwapNodes(t *testing.T) {
	testCases := []struct {
		got  []int
		gotK int
		want []int
	}{
		{
			got:  []int{1, 2, 3, 4, 5},
			gotK: 2,
			want: []int{1, 4, 3, 2, 5},
		},
		{
			got:  []int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5},
			gotK: 5,
			want: []int{7, 9, 6, 6, 8, 7, 3, 0, 9, 5},
		},
	}

	for _, testCase := range testCases {
		actual := swapNodes(FromSlice(testCase.got), testCase.gotK).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
