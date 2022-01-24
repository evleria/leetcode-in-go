package _832_flipping_an_image

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFlipAndInvertImage(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got:  [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want: [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{
			got:  [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			want: [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}

	for _, testCase := range testCases {
		actual := flipAndInvertImage(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
