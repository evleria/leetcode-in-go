package _953_verifying_an_alien_dictionary

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsAlienSorted(t *testing.T) {
	testCases := []struct {
		gotWords []string
		gotOrder string
		want     bool
	}{
		{
			gotWords: []string{"hello", "leetcode"},
			gotOrder: "hlabcdefgijkmnopqrstuvwxyz",
			want:     true,
		},
		{
			gotWords: []string{"word", "world", "row"},
			gotOrder: "worldabcefghijkmnpqstuvxyz",
			want:     false,
		},
	}

	for _, testCase := range testCases {
		actual := isAlienSorted(testCase.gotWords, testCase.gotOrder)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
