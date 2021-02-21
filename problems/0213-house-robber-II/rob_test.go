package _213_house_robber_II

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
			gotNums: []int{2, 3, 2},
			want:    3,
		},
		{
			gotNums: []int{1, 2, 3, 3},
			want:    5,
		},
	}

	for _, testCase := range testCases {
		actual := rob(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
