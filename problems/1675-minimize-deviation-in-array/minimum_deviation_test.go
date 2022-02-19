package _675_minimize_deviation_in_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumDeviation(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3, 4},
			want:    1,
		},
		{
			gotNums: []int{4, 1, 5, 20, 3},
			want:    3,
		},
		{
			gotNums: []int{2, 10, 8},
			want:    3,
		},
	}

	for _, testCase := range testCases {
		actual := minimumDeviation(testCase.gotNums)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
