package _520_detect_capital

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDetectCapitalUse(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "USA",
			want: true,
		},
		{
			got:  "FlaG",
			want: false,
		},
		{
			got:  "leetcode",
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := detectCapitalUse(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
