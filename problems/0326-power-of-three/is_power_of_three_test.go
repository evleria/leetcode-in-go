package _326_power_of_three

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPowerOfThree(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  27,
			want: true,
		},
		{
			got:  9,
			want: true,
		},
		{
			got:  0,
			want: false,
		},
		{
			got:  45,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isPowerOfThree(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
