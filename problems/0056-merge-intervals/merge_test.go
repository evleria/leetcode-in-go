package _056_merge_intervals

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got: [][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			want: [][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			got: [][]int{
				{1, 4},
				{4, 5},
			},
			want: [][]int{
				{1, 5},
			},
		},
		{
			got: [][]int{
				{1, 4},
				{0, 4},
			},
			want: [][]int{
				{0, 4},
			},
		},
	}

	for _, testCase := range testCases {
		actual := merge(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
