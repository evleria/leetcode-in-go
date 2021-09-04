package _834_sum_of_distances_in_tree

import (
	"fmt"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestSumOfDistancesInTree(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want []int
	}{
		{
			got:  [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			want: []int{8, 12, 6, 10, 10, 10},
		},
		{
			got:  [][]int{},
			want: []int{0},
		},
		{
			got:  [][]int{{1, 0}},
			want: []int{1, 1},
		},
	}

	for _, testCase := range testCases {
		actual := sumOfDistancesInTree(len(testCase.got)+1, testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
