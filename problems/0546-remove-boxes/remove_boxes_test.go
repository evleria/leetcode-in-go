package _546_remove_boxes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRemoveBoxes(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 3, 2, 2, 2, 3, 4, 3, 1},
			want: 23,
		},
		{
			got:  []int{1, 1, 1},
			want: 9,
		},
		{
			got:  []int{1},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := removeBoxes(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
