package _231_power_of_two

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPowerOfTwo(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  1,
			want: true,
		},
		{
			got:  16,
			want: true,
		},
		{
			got:  3,
			want: false,
		},
		{
			got:  5,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isPowerOfTwo(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
