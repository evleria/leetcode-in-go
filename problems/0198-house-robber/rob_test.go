package _198_house_robber

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRob(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3, 1},
			want:    4,
		},
		{
			gotNums: []int{2, 7, 9, 3, 1},
			want:    12,
		},
	}

	for _, testCase := range testCases {
		actual := rob(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
