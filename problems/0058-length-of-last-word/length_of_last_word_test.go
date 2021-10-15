package _058_length_of_last_word

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "Hello World",
			want: 5,
		},
		{
			got:  "   fly me   to   the moon  ",
			want: 4,
		},
		{
			got:  "luffy is still joyboy",
			want: 6,
		},
	}

	for _, testCase := range testCases {
		actual := lengthOfLastWord(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
