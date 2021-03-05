package _637_average_of_levels_in_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAverageOfLevels(t *testing.T) {
	testCases := []struct {
		got  []int
		want []float64
	}{
		{
			got:  []int{3, 9, 20, 15, 7},
			want: []float64{3, 14.5, 11},
		},
	}

	for _, testCase := range testCases {
		actual := averageOfLevels(FromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
