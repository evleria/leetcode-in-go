package _041_robot_bounded_in_circle

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsRobotBounded(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "GLRLLGLL",
			want: true,
		},
		{
			got:  "GGLLGG",
			want: true,
		},
		{
			got:  "GG",
			want: false,
		},
		{
			got:  "GL",
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := isRobotBounded(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
