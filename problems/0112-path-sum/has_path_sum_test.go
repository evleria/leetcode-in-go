package _112_path_sum

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHasPathSum(t *testing.T) {
	testCases := []struct {
		gotTree []int
		gotSum  int
		want    bool
	}{
		{
			gotTree: []int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, NULL, 1},
			gotSum:  22,
			want:    true,
		},
		{
			gotTree: []int{},
			gotSum:  0,
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := hasPathSum(BinaryTreeFromSlice(testCase.gotTree), testCase.gotSum)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
