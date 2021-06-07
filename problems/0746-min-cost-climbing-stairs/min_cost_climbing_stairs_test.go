package _746_min_cost_climbing_stairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{10, 15, 20},
			want:    15,
		},
		{
			gotNums: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want:    6,
		},
	}

	for _, testCase := range testCases {
		actual := minCostClimbingStairs(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
