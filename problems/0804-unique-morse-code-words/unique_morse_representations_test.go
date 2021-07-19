package _804_unique_morse_code_words

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"gin", "zen", "gig", "msg"},
			want: 2,
		},
		{
			got:  []string{"a"},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := uniqueMorseRepresentations(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
