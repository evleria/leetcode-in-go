package _035_search_insert_position

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{1, 3, 5, 6},
			gotTarget: 5,
			want:      2,
		},
		{
			gotNums:   []int{1, 3, 5, 6},
			gotTarget: 2,
			want:      1,
		},
		{
			gotNums:   []int{1, 3, 5, 6},
			gotTarget: 7,
			want:      4,
		},
		{
			gotNums:   []int{1, 3, 5, 6},
			gotTarget: 0,
			want:      0,
		},
	}

	for _, testCase := range testCases {
		actual := searchInsert(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
