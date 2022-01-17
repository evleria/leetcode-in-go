package _114_maximum_number_of_words_found_in_sentences

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMostWordsFound(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			want: 6,
		},
		{
			got:  []string{"please wait", "continue to fight", "continue to win"},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := mostWordsFound(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
