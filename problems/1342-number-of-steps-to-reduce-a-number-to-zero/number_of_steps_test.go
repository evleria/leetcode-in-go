package _342_number_of_steps_to_reduce_a_number_to_zero

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumberOfSteps(t *testing.T) {
	testCases := []struct {
		gotNum int
		want   int
	}{
		{
			gotNum: 14,
			want:   6,
		},
		{
			gotNum: 8,
			want:   4,
		},
		{
			gotNum: 123,
			want:   12,
		},
	}

	for _, testCase := range testCases {
		actual := numberOfSteps(testCase.gotNum)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
