package _202_happy_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsHappy(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  19,
			want: true,
		},
		{
			got:  2,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isHappy(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
