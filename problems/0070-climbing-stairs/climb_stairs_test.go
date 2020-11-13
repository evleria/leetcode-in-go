package _070_climbing_stairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  2,
			want: 2,
		},
		{
			got:  3,
			want: 3,
		},
		{
			got:  4,
			want: 5,
		},
		{
			got:  5,
			want: 8,
		},
	}

	for _, testCase := range testCases {
		actual := climbStairs(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
