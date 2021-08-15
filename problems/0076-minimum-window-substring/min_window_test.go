package _076_minimum_window_substring

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinWindow(t *testing.T) {
	testCases := []struct {
		gotS string
		gotT string
		want string
	}{
		{
			gotS: "ADOBECODEBANC",
			gotT: "ABC",
			want: "BANC",
		},
		{
			gotS: "a",
			gotT: "aa",
			want: "",
		},
		{
			gotS: "a",
			gotT: "a",
			want: "a",
		},
	}

	for _, testCase := range testCases {
		actual := minWindow(testCase.gotS, testCase.gotT)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
