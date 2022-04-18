package _230_kth_smallest_element_in_a_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKthSmallest(t *testing.T) {
	testCases := []struct {
		got  []int
		gotK int
		want int
	}{
		{
			got:  []int{3, 1, 4, NULL, 2},
			gotK: 1,
			want: 1,
		},
		{
			got:  []int{5, 3, 6, 2, 4, NULL, NULL, 1},
			gotK: 3,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := kthSmallest(BinaryTreeFromSlice(testCase.got), testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
