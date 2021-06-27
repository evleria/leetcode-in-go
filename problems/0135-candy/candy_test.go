package _135_candy

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCandy(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{2},
			want:    1,
		},
		{
			gotNums: []int{1, 1, 1},
			want:    3,
		},
		{
			gotNums: []int{1, 2, 3},
			want:    6,
		},
		{
			gotNums: []int{1, 2, 3, 3},
			want:    7,
		},
		{
			gotNums: []int{3, 2, 1},
			want:    6,
		},
		{
			gotNums: []int{1, 2, 3, 4, 5, 3, 2, 1, 2, 6, 5, 4, 3, 3, 2, 1, 1, 3, 3, 3, 4, 2},
			want:    47,
		},
		{
			gotNums: []int{1, 6, 10, 8, 7, 3, 2},
			want:    18,
		},
	}

	for _, testCase := range testCases {
		actual := candy(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotNums)
	}
}
