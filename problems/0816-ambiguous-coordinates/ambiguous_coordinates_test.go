package _816_ambiguous_coordinates

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAmbiguousCoordinates(t *testing.T) {
	testCases := []struct {
		gotS string
		want []string
	}{
		{
			gotS: "(123)",
			want: []string{"(1, 23)", "(1, 2.3)", "(12, 3)", "(1.2, 3)"},
		},
		{
			gotS: "(00011)",
			want: []string{"(0, 0.011)", "(0.001, 1)"},
		},
		{
			gotS: "(01200)",
			want: []string{"(0, 1200)", "(0.1, 200)"},
		},
	}

	for _, testCase := range testCases {
		actual := ambiguousCoordinates(testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
