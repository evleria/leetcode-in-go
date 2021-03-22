package _966_vowel_spellchecker

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSpellChecker(t *testing.T) {
	testCases := []struct {
		gotList    []string
		gotQueries []string
		want       []string
	}{
		{
			gotList:    []string{"KiTe", "kite", "hare", "Hare"},
			gotQueries: []string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			want:       []string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
	}

	for _, testCase := range testCases {
		actual := spellchecker(testCase.gotList, testCase.gotQueries)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
