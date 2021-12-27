package _103_rings_and_rods

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountPoints(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "B0B6G0R6R0R6G9",
			want: 1,
		},
		{
			got:  "B0R0G0R9R0B0G0",
			want: 1,
		},
		{
			got:  "G4",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := countPoints(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
