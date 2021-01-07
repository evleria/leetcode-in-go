package _024_swap_nodes_in_pairs

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSwapPairs(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
		{
			got:  []int{1},
			want: []int{1},
		},
		{
			got:  []int{},
			want: []int{},
		},
	}

	for _, testCase := range testCases {
		actual := swapPairs(FromSlice(testCase.got)).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
