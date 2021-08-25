package _633_sum_of_square_numbers

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestJudgeSquareSum(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    bool
	}{
		{
			gotNums: 5,
			want:    true,
		},
		{
			gotNums: 3,
			want:    false,
		},
		{
			gotNums: 4,
			want:    true,
		},
		{
			gotNums: 1,
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := judgeSquareSum(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
