package _139_word_break

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWordBreak(t *testing.T) {
	testCases := []struct {
		gotS    string
		gotWord []string
		want    bool
	}{
		{
			gotS: "leetcode",
			gotWord: []string{
				"leet", "code",
			},
			want: true,
		},
		{
			gotS: "applepenapple",
			gotWord: []string{
				"apple", "pen",
			},
			want: true,
		},
		{
			gotS: "catsandog",
			gotWord: []string{
				"cats", "dog", "sand", "and", "cat",
			},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := wordBreak(testCase.gotS, testCase.gotWord)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
