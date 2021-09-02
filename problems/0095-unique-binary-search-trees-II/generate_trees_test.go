package _095_unique_binary_search_trees_II

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGenerateTrees(t *testing.T) {
	testCases := []struct {
		got  int
		want [][]int
	}{
		{
			got: 3,
			want: [][]int{
				{1, NULL, 2, NULL, 3},
				{1, NULL, 3, 2},
				{2, 1, 3},
				{3, 1, NULL, NULL, 2}, {3, 2, NULL, 1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := generateTrees(testCase.got)
		slices := make([][]int, 0, len(actual))
		for _, tree := range actual {
			slices = append(slices, tree.ToSlice())
		}

		assert.Check(t, is.DeepEqual(slices, testCase.want))
	}
}
