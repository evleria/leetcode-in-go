package _465_maximum_area_of_a_piece_of_cake_after_horizontal_and_vertical_cuts

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		gotH          int
		gotW          int
		gotHorizontal []int
		gotVertical   []int
		want          int
	}{
		{
			gotH:          5,
			gotW:          4,
			gotHorizontal: []int{1, 2, 4},
			gotVertical:   []int{1, 3},
			want:          4,
		},
		{
			gotH:          5,
			gotW:          4,
			gotHorizontal: []int{3, 1},
			gotVertical:   []int{1},
			want:          6,
		},
		{
			gotH:          5,
			gotW:          4,
			gotHorizontal: []int{3},
			gotVertical:   []int{3},
			want:          9,
		},
	}

	for _, testCase := range testCases {
		actual := maxArea(testCase.gotH, testCase.gotW, testCase.gotHorizontal, testCase.gotVertical)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
