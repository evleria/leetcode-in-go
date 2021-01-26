package _290_word_pattern

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWordPattern(t *testing.T) {
	testCases := []struct {
		gotPattern string
		gotS       string
		want       bool
	}{
		{
			gotPattern: "abba",
			gotS:       "dog cat cat dog",
			want:       true,
		},
		{
			gotPattern: "abba",
			gotS:       "dog cat cat fish",
			want:       false,
		},
		{
			gotPattern: "abba",
			gotS:       "dog dog dog dog",
			want:       false,
		},
	}

	for _, testCase := range testCases {
		actual := wordPattern(testCase.gotPattern, testCase.gotS)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
