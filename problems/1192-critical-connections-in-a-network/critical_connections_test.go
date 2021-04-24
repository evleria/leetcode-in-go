package _192_critical_connections_in_a_network

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCriticalConnections(t *testing.T) {
	testCases := []struct {
		gotN int
		got  [][]int
		want [][]int
	}{
		{
			gotN: 4,
			got: [][]int{
				{0, 1},
				{1, 2},
				{2, 0},
				{1, 3},
			},
			want: [][]int{
				{1, 3},
			},
		},
	}

	for _, testCase := range testCases {
		actual := criticalConnections(testCase.gotN, testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
