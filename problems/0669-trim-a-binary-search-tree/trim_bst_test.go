package _669_trim_a_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTrimBst(t *testing.T) {
	testCases := []struct {
		gotRoot []int
		gotLow  int
		gotHigh int
		want    []int
	}{
		{
			gotRoot: []int{1, 0, 2},
			gotLow:  1,
			gotHigh: 2,
			want:    []int{1, NULL, 2},
		},
		{
			gotRoot: []int{3, 0, 4, NULL, 2, NULL, NULL, 1},
			gotLow:  1,
			gotHigh: 3,
			want:    []int{3, 2, NULL, 1},
		},
	}

	for _, testCase := range testCases {
		actual := trimBST(BinaryTreeFromSlice(testCase.gotRoot), testCase.gotLow, testCase.gotHigh).ToSlice()

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
