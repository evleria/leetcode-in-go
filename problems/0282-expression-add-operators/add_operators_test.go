package _282_expression_add_operators

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddOperators(t *testing.T) {
	testCases := []struct {
		gotNum    string
		gotTarget int
		want      []string
	}{
		{
			gotNum:    "232",
			gotTarget: 8,
			want:      []string{"2+3*2", "2*3+2"},
		},
		{
			gotNum:    "123",
			gotTarget: 6,
			want:      []string{"1+2+3", "1*2*3"},
		},
		{
			gotNum:    "105",
			gotTarget: 5,
			want:      []string{"1*0+5", "10-5"},
		},
		{
			gotNum:    "00",
			gotTarget: 0,
			want:      []string{"0+0", "0-0", "0*0"},
		},
	}

	for _, testCase := range testCases {
		actual := addOperators(testCase.gotNum, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
