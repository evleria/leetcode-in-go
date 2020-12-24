package _048_rotate_image

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for _, testCase := range testCases {
		rotate(testCase.got)

		assert.Check(t, is.DeepEqual(testCase.got, testCase.want))
	}
}
