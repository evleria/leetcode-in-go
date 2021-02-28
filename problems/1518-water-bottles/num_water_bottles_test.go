package _518_water_bottles

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumWaterBottles(t *testing.T) {
	testCases := []struct {
		gotB int
		gotE int
		want int
	}{
		{
			gotB: 9,
			gotE: 3,
			want: 13,
		},
		{
			gotB: 15,
			gotE: 4,
			want: 19,
		},
		{
			gotB: 2,
			gotE: 3,
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := numWaterBottles(testCase.gotB, testCase.gotE)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
