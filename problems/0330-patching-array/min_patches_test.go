package _330_patching_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinPatches(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotN    int
		want    int
	}{
		{
			gotNums: []int{1, 3},
			gotN:    6,
			want:    1,
		},
		{
			gotNums: []int{1, 5, 10},
			gotN:    20,
			want:    2,
		},
		{
			gotNums: []int{1, 2, 2},
			gotN:    5,
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := minPatches(testCase.gotNums, testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotNums)
	}
}
