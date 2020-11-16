package _111_minimum_depth_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinDepth(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 9, 20, NULL, NULL, 15, 7},
			want: 2,
		},
		{
			got:  []int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6},
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := minDepth(FromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
