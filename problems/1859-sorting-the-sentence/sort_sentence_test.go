package _859_sorting_the_sentence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortSentence(t *testing.T) {
	testCases := []struct {
		gotS string
		want string
	}{
		{
			gotS: "is2 sentence4 This1 a3",
			want: "This is a sentence",
		},
		{
			gotS: "Myself2 Me1 I4 and3",
			want: "Me Myself and I",
		},
	}

	for _, testCase := range testCases {
		actual := sortSentence(testCase.gotS)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotS)
	}
}
