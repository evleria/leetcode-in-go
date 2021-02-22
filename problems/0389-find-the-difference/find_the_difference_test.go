package _389_find_the_difference

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindTheDifference(t *testing.T) {
	testCases := []struct {
		gotS string
		gotT string
		want byte
	}{
		{
			gotS: "abcd",
			gotT: "caebd",
			want: 'e',
		},
		{
			gotS: "",
			gotT: "y",
			want: 'y',
		},
		{
			gotS: "a",
			gotT: "aa",
			want: 'a',
		},
	}

	for _, testCase := range testCases {
		actual := findTheDifference(testCase.gotS, testCase.gotT)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
