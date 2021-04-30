package _970_powerful_integers

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPowerfulIntegers(t *testing.T) {
	testCases := []struct {
		gotX     int
		gotY     int
		gotBound int
		want     []int
	}{
		{
			gotX:     2,
			gotY:     3,
			gotBound: 10,
			want:     []int{2, 3, 4, 5, 7, 9, 10},
		},
		{
			gotX:     3,
			gotY:     5,
			gotBound: 15,
			want:     []int{2, 4, 6, 8, 10, 14},
		},
		{
			gotX:     1,
			gotY:     2,
			gotBound: 10,
			want:     []int{2, 3, 5, 9},
		},
	}

	for _, testCase := range testCases {
		actual := powerfulIntegers(testCase.gotX, testCase.gotY, testCase.gotBound)
		sort.Ints(actual)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
