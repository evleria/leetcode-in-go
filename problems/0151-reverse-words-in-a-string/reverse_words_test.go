package _151_reverse_words_in_a_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "the sky is blue",
			want: "blue is sky the",
		},
		{
			got:  "  hello world  ",
			want: "world hello",
		},
		{
			got:  "a good   example",
			want: "example good a",
		},
		{
			got:  "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			got:  "Alice does not even like bob",
			want: "bob like even not does Alice",
		},
	}

	for _, testCase := range testCases {
		actual := reverseWords(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
