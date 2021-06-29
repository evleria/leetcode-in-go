package _836_rectangle_overlap

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsRectangleOverlap(t *testing.T) {
	testCases := []struct {
		gotRec1 []int
		gotRec2 []int
		want    bool
	}{
		{
			gotRec1: []int{0, 0, 2, 2},
			gotRec2: []int{1, 1, 3, 3},
			want:    true,
		},
		{
			gotRec1: []int{0, 0, 1, 1},
			gotRec2: []int{1, 0, 2, 1},
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := isRectangleOverlap(testCase.gotRec1, testCase.gotRec2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
