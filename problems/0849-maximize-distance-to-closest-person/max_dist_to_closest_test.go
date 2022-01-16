package _849_maximize_distance_to_closest_person

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxDistToClosest(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 0, 0, 0, 1, 0, 1},
			want: 2,
		},
		{
			got:  []int{1, 0, 0, 0},
			want: 3,
		},
		{
			got:  []int{0, 0, 0, 1},
			want: 3,
		},
		{
			got:  []int{0, 1},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := maxDistToClosest(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
