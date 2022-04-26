package _584_min_cost_to_connect_all_points

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinCostConnectPoints(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			want:    20,
		},
		{
			gotNums: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			want:    18,
		},
	}

	for _, testCase := range testCases {
		actual := minCostConnectPoints(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
