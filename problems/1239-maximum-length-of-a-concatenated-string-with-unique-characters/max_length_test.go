package _239_maximum_length_of_a_concatenated_string_with_unique_characters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxLenght(t *testing.T) {
	testCases := []struct {
		gotS []string
		want int
	}{
		{
			gotS: []string{"un", "iq", "ue"},
			want: 4,
		},
		{
			gotS: []string{"cha", "r", "act", "ers"},
			want: 6,
		},
		{
			gotS: []string{"abcdefghijklmnopqrstuvwxyz"},
			want: 26,
		},
	}

	for _, testCase := range testCases {
		actual := maxLength(testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
