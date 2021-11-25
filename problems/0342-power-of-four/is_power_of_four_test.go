package _342_power_of_four

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPowerOfFour(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  16,
			want: true,
		},
		{
			got:  1,
			want: true,
		},
		{
			got:  5,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isPowerOfFour(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
