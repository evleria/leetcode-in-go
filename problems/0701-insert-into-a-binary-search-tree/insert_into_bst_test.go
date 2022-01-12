package _701_insert_into_a_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestInsertIntoBST(t *testing.T) {
	testCases := []struct {
		got    []int
		gotVal int
		want   []int
	}{
		{
			got:    []int{4, 2, 7, 1, 3},
			gotVal: 5,
			want:   []int{4, 2, 7, 1, 3, 5},
		},
		{
			got:    []int{40, 20, 60, 10, 30, 50, 70},
			gotVal: 25,
			want:   []int{40, 20, 60, 10, 30, 50, 70, NULL, NULL, 25},
		},
		{
			got:    []int{4, 2, 7, 1, 3, NULL, NULL, NULL, NULL, NULL, NULL},
			gotVal: 5,
			want:   []int{4, 2, 7, 1, 3, 5},
		},
	}

	for _, testCase := range testCases {
		actual := insertIntoBST(BinaryTreeFromSlice(testCase.got), testCase.gotVal).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
