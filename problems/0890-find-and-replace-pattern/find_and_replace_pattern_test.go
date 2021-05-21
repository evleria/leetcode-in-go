package _890_find_and_replace_pattern

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFindAndReplacePattern(t *testing.T) {
	testCases := []struct {
		gotW []string
		gotR string
		want []string
	}{
		{
			gotW: []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
			gotR: "abb",
			want: []string{"mee", "aqq"},
		},
		{
			gotW: []string{"a", "b", "c"},
			gotR: "a",
			want: []string{"a", "b", "c"},
		},
	}

	for _, testCase := range testCases {
		actual := findAndReplacePattern(testCase.gotW, testCase.gotR)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
