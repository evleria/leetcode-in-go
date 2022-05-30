package _029_divide_two_integers

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		gotDividend int
		gotDivisior int
		want        int
	}{
		{
			gotDividend: 10,
			gotDivisior: 3,
			want:        3,
		},
		{
			gotDividend: 7,
			gotDivisior: -3,
			want:        -2,
		},
	}

	for _, testCase := range testCases {
		actual := divide(testCase.gotDividend, testCase.gotDivisior)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
