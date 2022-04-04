package _030_matrix_cells_in_distance_order

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAllCellsDistOrder(t *testing.T) {
	testCases := []struct {
		gotRows    int
		gotCols    int
		gotRCenter int
		gotCCenter int
		want       [][]int
	}{
		{
			gotRows:    1,
			gotCols:    2,
			gotRCenter: 0,
			gotCCenter: 0,
			want: [][]int{
				{0, 0},
				{0, 1},
			},
		},
		{
			gotRows:    2,
			gotCols:    2,
			gotRCenter: 0,
			gotCCenter: 1,
			want: [][]int{
				{0, 1},
				{0, 0},
				{1, 1},
				{1, 0},
			},
		},
		{
			gotRows:    2,
			gotCols:    3,
			gotRCenter: 1,
			gotCCenter: 2,
			want: [][]int{
				{1, 2},
				{0, 2},
				{1, 1},
				{0, 1},
				{1, 0},
				{0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		actual := allCellsDistOrder(testCase.gotRows, testCase.gotCols, testCase.gotRCenter, testCase.gotCCenter)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotRows)
	}
}
