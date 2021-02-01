package _605_can_place_flowers

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanPlaceFlowers(t *testing.T) {
	testCases := []struct {
		gotF []int
		gotN int
		want bool
	}{
		{
			gotF: []int{1, 0, 0, 0, 1},
			gotN: 1,
			want: true,
		},
		{
			gotF: []int{1, 0, 0, 0, 1},
			gotN: 2,
			want: false,
		},
		{
			gotF: []int{1, 0, 1, 0, 1, 0, 1},
			gotN: 0,
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := canPlaceFlowers(testCase.gotF, testCase.gotN)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
