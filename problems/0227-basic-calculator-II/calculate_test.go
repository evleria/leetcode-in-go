package _227_basic_calculator_II

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "3+2*2",
			want: 7,
		},
		{
			got:  " 3/2 ",
			want: 1,
		},
		{
			got:  " 3+5 / 2 ",
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := calculate(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
