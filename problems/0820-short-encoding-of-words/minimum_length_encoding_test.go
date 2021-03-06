package _820_short_encoding_of_words

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumLengthEncoding(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"time", "me", "bell"},
			want: 10,
		},
		{
			got:  []string{"t"},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := minimumLengthEncoding(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
