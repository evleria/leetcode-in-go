package _075_sort_colors

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			got:  []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			got:  []int{1, 0, 0},
			want: []int{0, 0, 1},
		},
		{
			got:  []int{1, 2, 2},
			want: []int{1, 2, 2},
		},
		{
			got:  []int{2, 2, 1},
			want: []int{1, 2, 2},
		},
		{
			got:  []int{0},
			want: []int{0},
		},
	}

	for _, testCase := range testCases {
		sortColors(testCase.got)

		assert.Check(t, is.DeepEqual(testCase.got, testCase.want))
	}
}
