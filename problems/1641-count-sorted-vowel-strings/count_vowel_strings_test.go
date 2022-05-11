package _641_count_sorted_vowel_strings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountVowelStrings(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  1,
			want: 5,
		},
		{
			got:  2,
			want: 15,
		},
		{
			got:  33,
			want: 66045,
		},
	}

	for _, testCase := range testCases {
		actual := countVowelStrings(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
