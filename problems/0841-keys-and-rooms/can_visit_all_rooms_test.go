package _841_keys_and_rooms

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestCanVisitAllRooms(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    bool
	}{
		{
			gotNums: [][]int{{1}, {2}, {3}, {}},
			want:    true,
		},
		{
			gotNums: [][]int{{1, 3}, {3, 0, 1}, {2}, {0}},
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := canVisitAllRooms(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
