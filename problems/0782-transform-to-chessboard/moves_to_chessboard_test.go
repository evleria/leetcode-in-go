package _782_transform_to_chessboard

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMovesToChessboard(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{1, 0, 0, 1},
				{1, 0, 0, 1}},
			want: 2,
		},
		{
			got: [][]int{
				{0, 1},
				{1, 0},
			},
			want: 0,
		},
		{
			got: [][]int{
				{1, 0},
				{1, 0},
			},
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := movesToChessboard(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
