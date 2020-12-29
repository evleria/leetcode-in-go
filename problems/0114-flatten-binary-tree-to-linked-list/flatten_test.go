package _114_flatten_binary_tree_to_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFlatten(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 5, 3, 4, NULL, 6},
			want: []int{1, NULL, 2, NULL, 3, NULL, 4, NULL, 5, NULL, 6},
		},
	}

	for _, testCase := range testCases {
		tree := FromSlice(testCase.got)
		flatten(tree)

		assert.Check(t, is.DeepEqual(tree.ToSlice(), testCase.want), testCase.got)
	}
}
