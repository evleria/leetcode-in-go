package _189_maximum_number_of_balloons

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "nlaebolko",
			want: 1,
		},
		{
			got:  "loonbalxballpoon",
			want: 2,
		},
		{
			got:  "leetcode",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxNumberOfBalloons(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
