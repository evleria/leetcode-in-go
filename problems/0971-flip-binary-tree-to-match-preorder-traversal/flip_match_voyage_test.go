package _971_flip_binary_tree_to_match_preorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFlipMatchVoyage(t *testing.T) {
	testCases := []struct {
		gotTree   []int
		gotVoyage []int
		want      []int
	}{
		{
			gotTree:   []int{1, 2},
			gotVoyage: []int{2, 1},
			want:      []int{-1},
		},
		{
			gotTree:   []int{1, 2, 3},
			gotVoyage: []int{1, 3, 2},
			want:      []int{1},
		},
		{
			gotTree:   []int{1, 2, 3},
			gotVoyage: []int{1, 2, 3},
			want:      []int{},
		},
	}

	for _, testCase := range testCases {
		actual := flipMatchVoyage(FromSlice(testCase.gotTree), testCase.gotVoyage)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
