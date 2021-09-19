package _115_distinct_subsequences

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumDistinct(t *testing.T) {
	testCases := []struct {
		gotS string
		gotT string
		want int
	}{
		{
			gotS: "rabbbit",
			gotT: "rabbit",
			want: 3,
		},
		{
			gotS: "babgbag",
			gotT: "bag",
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := numDistinct(testCase.gotS, testCase.gotT)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
