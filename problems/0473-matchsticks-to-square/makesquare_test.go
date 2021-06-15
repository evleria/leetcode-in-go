package _473_matchsticks_to_square

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMakesSquare(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
			want: true,
		},
		{
			got:  []int{1, 1, 2, 2, 2},
			want: true,
		},
		{
			got:  []int{3, 3, 3, 3, 4},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := makesquare(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
