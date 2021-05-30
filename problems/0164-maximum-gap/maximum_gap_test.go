package _164_maximum_gap

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximumGap(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{3, 6, 9, 1},
			want:    3,
		},
		{
			gotNums: []int{10, 0},
			want:    10,
		},
		{
			gotNums: []int{10},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := maximumGap(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
