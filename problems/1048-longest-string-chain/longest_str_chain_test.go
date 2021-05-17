package _048_longest_string_chain

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestLongestStrChain(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got:  []string{"a", "b", "ba", "bca", "bda", "bdca"},
			want: 4,
		},
		{
			got:  []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"},
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := longestStrChain(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
