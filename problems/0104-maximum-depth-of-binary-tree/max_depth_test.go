package _104_maximum_depth_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{3, 9, 20, NULL, NULL, 15, 7},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := maxDepth(FromSlice(testCase.got))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
