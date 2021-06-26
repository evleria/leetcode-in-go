package _614_maximum_nesting_depth_of_the_parentheses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		gotS string
		want int
	}{
		{
			gotS: "(1+(2*3)+((8)/4))+1",
			want: 3,
		},
		{
			gotS: "(1)+((2))+(((3)))",
			want: 3,
		},
		{
			gotS: "1+(2*3)/(2-1)",
			want: 1,
		},
		{
			gotS: "1",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxDepth(testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
