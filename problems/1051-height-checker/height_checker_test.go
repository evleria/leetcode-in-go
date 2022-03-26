package _051_height_checker

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHeightChecker(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 1, 4, 2, 1, 3},
			want: 3,
		},
		{
			got:  []int{5, 1, 2, 3, 4},
			want: 5,
		},
		{
			got:  []int{1, 2, 3, 4, 5},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := heightChecker(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
