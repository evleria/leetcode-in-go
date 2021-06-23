package _938_range_sum_of_bst

import (
	"fmt"
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRangeSumBST(t *testing.T) {
	testCases := []struct {
		gotRoot []int
		gotL    int
		gotH    int
		want    int
	}{
		{
			gotRoot: []int{10, 5, 15, 3, 7, NULL, 18},
			gotL:    7,
			gotH:    15,
			want:    32,
		},
		{
			gotRoot: []int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6},
			gotL:    6,
			gotH:    10,
			want:    23,
		},
	}

	for _, testCase := range testCases {
		actual := rangeSumBST(BinaryTreeFromSlice(testCase.gotRoot), testCase.gotL, testCase.gotH)

		assert.Check(t, is.Equal(actual, testCase.want), fmt.Sprintf("%v", testCase))
	}
}
