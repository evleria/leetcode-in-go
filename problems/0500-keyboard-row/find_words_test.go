package _500_keyboard_row

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFindWords(t *testing.T) {
	testCases := []struct {
		got  []string
		want []string
	}{
		{
			got:  []string{"Hello", "Alaska", "Dad", "Peace"},
			want: []string{"Alaska", "Dad"},
		},
	}

	for _, testCase := range testCases {
		actual := findWords(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
