package _906_super_palindromes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSuperpalindromesInRange(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want int
	}{
		{
			got1: "4",
			got2: "1000",
			want: 4,
		},

		{
			got1: "1",
			got2: "2",
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := superpalindromesInRange(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
