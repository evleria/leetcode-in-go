package _823_binary_trees_with_factors

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumFactoredBinaryTrees(t *testing.T) {
	testCases := []struct {
		gotArr []int
		want   int
	}{
		{
			gotArr: []int{2, 4},
			want:   3,
		},
		{
			gotArr: []int{2, 4, 5, 10},
			want:   7,
		},
	}

	for _, testCase := range testCases {
		actual := numFactoredBinaryTrees(testCase.gotArr)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
