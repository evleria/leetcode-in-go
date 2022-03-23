package _991_broken_calculator

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBrokenCalc(t *testing.T) {
	testCases := []struct {
		gotN int
		gotK int
		want int
	}{
		{
			gotN: 5,
			gotK: 3,
			want: 2,
		},
		{
			gotN: 5,
			gotK: 8,
			want: 2,
		},
		{
			gotN: 3,
			gotK: 10,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := brokenCalc(testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
