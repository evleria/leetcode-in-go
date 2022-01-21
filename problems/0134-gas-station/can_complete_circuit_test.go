package _134_gas_station

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanCompleteCircuit(t *testing.T) {
	testCases := []struct {
		gotG []int
		gotC []int
		want int
	}{
		{
			gotG: []int{1, 2, 3, 4, 5},
			gotC: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			gotG: []int{2, 3, 4},
			gotC: []int{3, 4, 3},
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := canCompleteCircuit(testCase.gotG, testCase.gotC)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
