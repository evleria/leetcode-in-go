package _260_shift_2d_grid

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestShiftGrid(t *testing.T) {
	testCases := []struct {
		got  [][]int
		gotK int
		want [][]int
	}{
		{
			got:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			gotK: 1,
			want: [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			got:  [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			gotK: 4,
			want: [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			got:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			gotK: 9,
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}

	for _, testCase := range testCases {
		actual := shiftGrid(testCase.got, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
