package _221_split_a_string_in_balanced_strings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBalancedStringSplit(t *testing.T) {
	testCases := []struct {
		gotS string
		want int
	}{
		{
			gotS: "RLRRLLRLRL",
			want: 4,
		},
		{
			gotS: "RLLLLRRRLR",
			want: 3,
		},
		{
			gotS: "LLLLRRRR",
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := balancedStringSplit(testCase.gotS)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
