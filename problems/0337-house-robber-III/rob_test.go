package _337_house_robber_III

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRob(t *testing.T) {
	testCases := []struct {
		gotTree []int
		want    int
	}{
		{
			gotTree: []int{3, 2, 3, NULL, 3, NULL, 1},
			want:    7,
		},
		{
			gotTree: []int{3, 4, 5, 1, 3, NULL, 1},
			want:    9,
		},
	}

	for _, testCase := range testCases {
		tree := BinaryTreeFromSlice(testCase.gotTree)
		actual := rob(tree)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
