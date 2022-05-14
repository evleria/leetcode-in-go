package _743_network_delay_time

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNetworkDelayTime(t *testing.T) {
	testCases := []struct {
		gotTimes [][]int
		gotN     int
		gotK     int
		want     int
	}{
		{
			gotTimes: [][]int{
				{2, 1, 1},
				{2, 3, 1},
				{3, 4, 1},
			},
			gotN: 4,
			gotK: 2,
			want: 2,
		},
		{
			gotTimes: [][]int{{1, 2, 1}},
			gotN:     2,
			gotK:     1,
			want:     1,
		},
		{
			gotTimes: [][]int{{1, 2, 1}},
			gotN:     2,
			gotK:     2,
			want:     -1,
		},
	}

	for _, testCase := range testCases {
		actual := networkDelayTime(testCase.gotTimes, testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
