package _450_delete_node_in_a_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		got    []int
		gotKey int
		want   []int
	}{
		{
			got:    []int{5, 3, 6, 2, 4, NULL, 7},
			gotKey: 3,
			want:   []int{5, 4, 6, 2, NULL, NULL, 7},
		},
		{
			got:    []int{5, 3, 6, 2, 4, NULL, 7},
			gotKey: 0,
			want:   []int{5, 3, 6, 2, 4, NULL, 7},
		},
		{
			got:    []int{},
			gotKey: 0,
			want:   []int{},
		},
	}

	for _, testCase := range testCases {
		actual := deleteNode(BinaryTreeFromSlice(testCase.got), testCase.gotKey).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
