package _792_number_of_matching_subsequences

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumMatchingSubseq(t *testing.T) {
	testCases := []struct {
		gotS string
		gotW []string
		want int
	}{
		{
			gotS: "abcde",
			gotW: []string{"a", "bb", "acd", "ace"},
			want: 3,
		},
		{
			gotS: "dsahjpjauf",
			gotW: []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := numMatchingSubseq(testCase.gotS, testCase.gotW)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotS)
	}
}
