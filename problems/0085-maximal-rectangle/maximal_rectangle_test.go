package _085_maximal_rectangle

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximalRectangle(t *testing.T) {
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
			want: 6,
		},
		{
			got:  [][]byte{{'0'}},
			want: 0,
		},
		{
			got:  [][]byte{{'1'}},
			want: 1,
		},
		{
			got:  [][]byte{{'0', '0'}},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maximalRectangle(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
