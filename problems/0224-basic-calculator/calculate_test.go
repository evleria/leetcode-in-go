package _224_basic_calculator

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "1 + 1",
			want: 2,
		},
		{
			got:  "-5+10",
			want: 5,
		},
		{
			got:  "10-(2+3)",
			want: 5,
		},
		{
			got:  " 2-1 + 2 ",
			want: 3,
		},
		{
			got:  "(1+(4+5+2)-3)+(6+8)",
			want: 23,
		},
	}

	for _, testCase := range testCases {
		actual := calculate(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
