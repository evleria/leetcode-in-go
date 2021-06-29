package _67_valid_perfect_square

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsPerfectSquare(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  16,
			want: true,
		},
		{
			got:  14,
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isPerfectSquare(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
