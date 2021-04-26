package _642_furthest_building_you_can_reach

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFurthestBuilding(t *testing.T) {
	testCases := []struct {
		gotHeights []int
		gotBricks  int
		gotLadders int
		want       int
	}{
		{
			gotHeights: []int{1, 5, 1, 2, 3, 4, 10000},
			gotBricks:  4,
			gotLadders: 1,
			want:       5,
		},
		{
			gotHeights: []int{4, 2, 7, 6, 9, 14, 12},
			gotBricks:  5,
			gotLadders: 1,
			want:       4,
		},
		{
			gotHeights: []int{4, 12, 2, 7, 3, 18, 20, 3, 19},
			gotBricks:  10,
			gotLadders: 2,
			want:       7,
		},
		{
			gotHeights: []int{14, 3, 19, 3},
			gotBricks:  17,
			gotLadders: 0,
			want:       3,
		},
	}

	for _, testCase := range testCases {
		actual := furthestBuilding(testCase.gotHeights, testCase.gotBricks, testCase.gotLadders)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
