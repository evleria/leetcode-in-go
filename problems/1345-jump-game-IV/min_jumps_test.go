package _345_jump_game_IV

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinJumps(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404},
			want: 3,
		},
		{
			got:  []int{7},
			want: 0,
		},
		{
			got:  []int{7, 6, 9, 6, 9, 6, 9, 7},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := minJumps(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
