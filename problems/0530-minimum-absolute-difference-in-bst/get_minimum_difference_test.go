package _530_minimum_absolute_difference_in_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetMinimumDifference(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{4, 2, 6, 1, 3},
			want: 1,
		},
		{
			got:  []int{1, 0, 48, NULL, NULL, 12, 49},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := getMinimumDifference(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
