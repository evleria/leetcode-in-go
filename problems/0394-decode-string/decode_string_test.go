package _394_decode_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "3[a]2[bc]",
			want: "aaabcbc",
		},
		{
			got:  "3[a2[c]]",
			want: "accaccacc",
		},
		{
			got:  "2[abc]3[cd]ef",
			want: "abcabccdcdcdef",
		},
		{
			got:  "abc3[cd]xyz",
			want: "abccdcdcdxyz",
		},
	}

	for _, testCase := range testCases {
		actual := decodeString(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
