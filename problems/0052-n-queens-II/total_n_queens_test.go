package _052_n_queens_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTotalNQueens(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    int
	}{
		{
			gotNums: 4,
			want:    2,
		},
		{
			gotNums: 1,
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := totalNQueens(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
