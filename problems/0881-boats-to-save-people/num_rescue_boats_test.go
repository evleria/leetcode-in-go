package _881_boats_to_save_people

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumRescueBoats(t *testing.T) {
	testCases := []struct {
		gotPeople []int
		gotLimit  int
		want      int
	}{
		{
			gotPeople: []int{1, 2},
			gotLimit:  3,
			want:      1,
		},
		{
			gotPeople: []int{3, 2, 2, 1},
			gotLimit:  3,
			want:      3,
		},
		{
			gotPeople: []int{3, 5, 3, 4},
			gotLimit:  5,
			want:      4,
		},
	}

	for _, testCase := range testCases {
		actual := numRescueBoats(testCase.gotPeople, testCase.gotLimit)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
