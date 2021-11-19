package _657_robot_return_to_origin

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestJudgeCircle(t *testing.T) {
	testCases := []struct {
		got  string
		want bool
	}{
		{
			got:  "UD",
			want: true,
		},
		{
			got:  "LL",
			want: false,
		},
		{
			got:  "LDRRLRUULR",
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := judgeCircle(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
