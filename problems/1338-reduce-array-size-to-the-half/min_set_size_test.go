package _338_reduce_array_size_to_the_half

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinSetSize(t *testing.T) {
	testCases := []struct {
		gotArr []int
		want   int
	}{
		{
			gotArr: []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7},
			want:   2,
		},
		{
			gotArr: []int{7, 7, 7, 7, 7, 7},
			want:   1,
		},
		{
			gotArr: []int{1, 9},
			want:   1,
		},
		{
			gotArr: []int{1000, 1000, 3, 7},
			want:   1,
		},
		{
			gotArr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:   5,
		},
	}

	for _, testCase := range testCases {
		actual := minSetSize(testCase.gotArr)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
