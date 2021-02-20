package _739_daily_temperatures

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:    []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := dailyTemperatures(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
