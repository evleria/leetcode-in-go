package _029_two_city_scheduling

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTwoCitySchedCost(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}},
			want: 110,
		},
		{
			got:  [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}},
			want: 1859,
		},
	}

	for _, testCase := range testCases {
		actual := twoCitySchedCost(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
