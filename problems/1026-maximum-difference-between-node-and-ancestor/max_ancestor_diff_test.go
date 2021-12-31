package _026_maximum_difference_between_node_and_ancestor

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxAncestorDiff(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{8, 3, 10, 1, 6, NULL, 14, NULL, NULL, 4, 7, 13},
			want: 7,
		},
		{
			got:  []int{1, NULL, 2, NULL, 0, 3},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := maxAncestorDiff(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
