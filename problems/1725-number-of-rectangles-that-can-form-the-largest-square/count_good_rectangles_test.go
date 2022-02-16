package _725_number_of_rectangles_that_can_form_the_largest_square

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountGoodRectangles(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
			want: 3,
		},
		{
			got:  [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := countGoodRectangles(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
