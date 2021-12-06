package _217_minimum_cost_to_move_chips_to_the_same_position

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinCostToMoveChips(t *testing.T) {
	testCases := []struct {
		gotHead []int
		want    int
	}{
		{
			gotHead: []int{1, 2, 3},
			want:    1,
		},
		{
			gotHead: []int{2, 2, 2, 3, 3},
			want:    2,
		},
		{
			gotHead: []int{1, 1000000000},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := minCostToMoveChips(testCase.gotHead)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
