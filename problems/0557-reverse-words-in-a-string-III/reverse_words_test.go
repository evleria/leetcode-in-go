package _557_reverse_words_in_a_string_III

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
			got:  "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			got:  "God Ding",
			want: "doG gniD",
		},
	}

	for _, testCase := range testCases {
		actual := reverseWords(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
