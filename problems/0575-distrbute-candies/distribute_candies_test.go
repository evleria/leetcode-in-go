package _575_distrbute_candies

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDistributeCandies(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 1, 2, 2, 3, 3},
			want:    3,
		},
		{
			gotNums: []int{6, 6, 6, 6},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := distributeCandies(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
