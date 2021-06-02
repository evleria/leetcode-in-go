package _097_interleaving_string

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsInterleave(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		got3 string
		want bool
	}{
		{
			got1: "aacaac",
			got2: "aacaaeaac",
			got3: "aacaaeaaeaacaac",
			want: false,
		},
		{
			got1: "aaaa",
			got2: "aa",
			got3: "aaa",
			want: false,
		},
		{
			got1: "aabcc",
			got2: "dbbca",
			got3: "aadbbcbcac",
			want: true,
		},
		{
			got1: "aabcc",
			got2: "dbbca",
			got3: "aadbbbaccc",
			want: false,
		},
		{
			got1: "",
			got2: "",
			got3: "",
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := isInterleave(testCase.got1, testCase.got2, testCase.got3)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
