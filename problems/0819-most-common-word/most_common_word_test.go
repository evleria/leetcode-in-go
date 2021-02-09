package _819_most_common_word

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMostCommonWord(t *testing.T) {
	testCases := []struct {
		gotParagraph string
		gotBanned    []string
		want         string
	}{
		{
			gotParagraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			gotBanned:    []string{"hit"},
			want:         "ball",
		},
		{
			gotParagraph: "a, a, a, a, b,b,b,c, c",
			gotBanned:    []string{"a"},
			want:         "b",
		},
	}

	for _, testCase := range testCases {
		actual := mostCommonWord(testCase.gotParagraph, testCase.gotBanned)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
