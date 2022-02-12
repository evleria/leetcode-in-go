package _127_word_ladder

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLadderLength(t *testing.T) {
	testCases := []struct {
		gotBegin string
		gotEnd   string
		gotList  []string
		want     int
	}{
		{
			gotBegin: "hit",
			gotEnd:   "cog",
			gotList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:     5,
		},
		{
			gotBegin: "hit",
			gotEnd:   "cog",
			gotList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:     0,
		},
	}

	for _, testCase := range testCases {
		actual := ladderLength(testCase.gotBegin, testCase.gotEnd, testCase.gotList)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
