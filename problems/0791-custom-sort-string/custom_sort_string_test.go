package _791_custom_sort_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCustomSortString(t *testing.T) {
	testCases := []struct {
		gotOrder string
		gotStr   string
		want     string
	}{
		{
			gotOrder: "cba",
			gotStr:   "abcd",
			want:     "cbad",
		},
	}

	for _, testCase := range testCases {
		actual := customSortString(testCase.gotOrder, testCase.gotStr)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotOrder)
	}
}
