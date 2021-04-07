package _704_determine_if_string_halves_are_alike

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHalvesAreAlike(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "book",
			want: true,
		},
		{
			got:  "textbook",
			want: false,
		},
		{
			got:  "AbCdEfGh",
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := halvesAreAlike(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
