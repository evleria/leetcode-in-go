package _051_n_queens

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	testCases := []struct {
		got  int
		want [][]string
	}{
		{
			got: 4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
		{
			got: 1,
			want: [][]string{
				{"Q"},
			},
		},
	}

	for _, testCase := range testCases {
		actual := solveNQueens(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
