package _684_count_the_number_of_consistent_strings

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCountConsistentStrings(t *testing.T) {
	testCases := []struct {
		gotAllowed string
		gotWords   []string
		want       int
	}{
		{
			gotAllowed: "ab",
			gotWords:   []string{"ad", "bd", "aaab", "baa", "badab"},
			want:       2,
		},
		{
			gotAllowed: "abc",
			gotWords:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			want:       7,
		},
	}

	for _, testCase := range testCases {
		actual := countConsistentStrings(testCase.gotAllowed, testCase.gotWords)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
