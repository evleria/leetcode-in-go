package _844_backspace_string_compare

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBackspaceCompare(t *testing.T) {
	testCases := []struct {
		gotS string
		gotT string
		want bool
	}{
		{
			gotS: "ab#c",
			gotT: "ad#c",
			want: true,
		},
		{
			gotS: "ab##",
			gotT: "c#d#",
			want: true,
		},
		{
			gotS: "a#c",
			gotT: "b",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := backspaceCompare(testCase.gotS, testCase.gotT)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
