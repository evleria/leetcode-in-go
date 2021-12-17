package _221_maximal_square

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximalSquare(t *testing.T) {
	testCases := []struct {
		got  [][]byte
		want int
	}{
		{
			got: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'}},
			want: 4,
		},
		{
			got:  [][]byte{{'0', '1'}, {'1', '0'}},
			want: 1,
		},
		{
			got:  [][]byte{{'0'}},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maximalSquare(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
