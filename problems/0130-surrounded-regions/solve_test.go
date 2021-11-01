package _130_surrounded_regions

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSolve(t *testing.T) {
	testCases := []struct {
		got  [][]byte
		want [][]byte
	}{
		{
			got: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			got: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
	}

	for _, testCase := range testCases {
		solve(testCase.got)

		assert.Check(t, is.DeepEqual(testCase.got, testCase.want))
	}
}
