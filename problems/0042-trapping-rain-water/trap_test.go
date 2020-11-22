package _042_trapping_rain_water

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTrap(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want: 6,
		},
		{
			got:  []int{4, 2, 0, 3, 2, 5},
			want: 9,
		},
	}

	for _, testCase := range testCases {
		actual := trap(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
