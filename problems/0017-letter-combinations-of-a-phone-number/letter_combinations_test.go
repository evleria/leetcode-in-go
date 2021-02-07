package _017_letter_combinations_of_a_phone_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLetterCombinations(t *testing.T) {
	testCases := []struct {
		got  string
		want []string
	}{
		{
			got:  "",
			want: []string{},
		},
		{
			got:  "2",
			want: []string{"a", "b", "c"},
		},
		{
			got:  "23",
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for _, testCase := range testCases {
		actual := letterCombinations(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
