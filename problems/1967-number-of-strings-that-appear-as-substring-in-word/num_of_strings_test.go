package _967_number_of_strings_that_appear_as_substring_in_word

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReversePrefix(t *testing.T) {
	testCases := []struct {
		gotPatterns []string
		gotWord     string
		want        int
	}{
		{
			gotPatterns: []string{"a", "abc", "bc", "d"},
			gotWord:     "abc",
			want:        3,
		},
		{
			gotPatterns: []string{"a", "b", "c"},
			gotWord:     "aaaaabbbbb",
			want:        2,
		},
		{
			gotPatterns: []string{"a", "a", "a"},
			gotWord:     "ab",
			want:        3,
		},
	}

	for _, testCase := range testCases {
		actual := numOfStrings(testCase.gotPatterns, testCase.gotWord)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
