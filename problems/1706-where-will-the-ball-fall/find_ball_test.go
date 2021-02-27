package _706_where_will_the_ball_fall

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindBall(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want []int
	}{
		{
			got: [][]int{
				{1, 1, 1, -1, -1},
				{1, 1, 1, -1, -1},
				{-1, -1, -1, 1, 1},
				{1, 1, 1, 1, -1},
				{-1, -1, -1, -1, -1},
			},
			want: []int{1, -1, -1, -1, -1},
		},
	}

	for _, testCase := range testCases {
		actual := findBall(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
