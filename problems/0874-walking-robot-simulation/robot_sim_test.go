package _874_walking_robot_simulation

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRobotSlim(t *testing.T) {
	testCases := []struct {
		gotCommands  []int
		gotObstacles [][]int
		want         int
	}{
		{
			gotCommands:  []int{4, -1, 3},
			gotObstacles: [][]int{},
			want:         25,
		},
		{
			gotCommands: []int{4, -1, 4, -2, 4},
			gotObstacles: [][]int{
				{2, 4},
			},
			want: 65,
		},
		{
			gotCommands: []int{-2, 8, 3, 7, -1},
			gotObstacles: [][]int{
				{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3},
			},
			want: 324,
		},
	}

	for _, testCase := range testCases {
		actual := robotSim(testCase.gotCommands, testCase.gotObstacles)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
