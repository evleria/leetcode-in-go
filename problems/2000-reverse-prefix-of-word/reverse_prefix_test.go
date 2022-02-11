package _000_reverse_prefix_of_word

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReversePrefix(t *testing.T) {
	testCases := []struct {
		gotWord string
		gotCh   byte
		want    string
	}{
		{
			gotWord: "abcdefd",
			gotCh:   'd',
			want:    "dcbaefd",
		},
		{
			gotWord: "xyxzxe",
			gotCh:   'z',
			want:    "zxyxxe",
		},
		{
			gotWord: "abcd",
			gotCh:   'z',
			want:    "abcd",
		},
	}

	for _, testCase := range testCases {
		actual := reversePrefix(testCase.gotWord, testCase.gotCh)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
