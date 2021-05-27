package _318_maximum_product_of_word_lengths

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxProduct(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			want: 16,
		},
		{
			got:  []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
			want: 4,
		},
		{
			got:  []string{"a", "aa", "aaa", "aaaa"},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxProduct(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
