package _113_path_sum_II

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPathSum(t *testing.T) {
	testCases := []struct {
		gotRoot   []int
		gotTarget int
		want      [][]int
	}{
		{
			gotRoot:   []int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1},
			gotTarget: 22,
			want:      [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			gotRoot:   []int{1, 2, 3},
			gotTarget: 5,
			want:      [][]int{},
		},
		{
			gotRoot:   []int{1, 2},
			gotTarget: 0,
			want:      [][]int{},
		},
	}

	for _, testCase := range testCases {
		actual := pathSum(BinaryTreeFromSlice(testCase.gotRoot), testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
