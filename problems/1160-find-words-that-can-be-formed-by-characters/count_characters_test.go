package _160_find_words_that_can_be_formed_by_characters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountCharacters(t *testing.T) {
	testCases := []struct {
		gotWords []string
		gotChars string
		want     int
	}{
		{
			gotWords: []string{"cat", "bt", "hat", "tree"},
			gotChars: "atach",
			want:     6,
		},
	}

	for _, testCase := range testCases {
		actual := countCharacters(testCase.gotWords, testCase.gotChars)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
