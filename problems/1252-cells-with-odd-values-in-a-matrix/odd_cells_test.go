package _252_cells_with_odd_values_in_a_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOddCells(t *testing.T) {
	testCases := []struct {
		gotM int
		gotN int
		gotI [][]int
		want int
	}{
		{
			gotM: 2,
			gotN: 3,
			gotI: [][]int{{0, 1}, {1, 1}},
			want: 6,
		},
		{
			gotM: 2,
			gotN: 2,
			gotI: [][]int{{1, 1}, {0, 0}},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := oddCells(testCase.gotM, testCase.gotN, testCase.gotI)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
