package _838_push_dominoes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPushDominoes(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "RR.L",
			want: "RR.L",
		},
		{
			got:  ".L.R...LR..L..",
			want: "LL.RR.LLRRLL..",
		},
		{
			got:  "R.R.L",
			want: "RRR.L",
		},
	}

	for _, testCase := range testCases {
		actual := pushDominoes(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
