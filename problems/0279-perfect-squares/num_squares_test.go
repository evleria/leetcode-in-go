package _279_perfect_squares

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumSquares(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  12,
			want: 3,
		},
		{
			got:  13,
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := numSquares(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
