package _423_maximum_points_you_can_obtain_from_gards

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxScore(t *testing.T) {
	testCases := []struct {
		gotG []int
		gotK int
		want int
	}{
		{
			gotG: []int{1, 2, 3, 4, 5, 6, 1},
			gotK: 3,
			want: 12,
		},
		{
			gotG: []int{2, 2, 2},
			gotK: 2,
			want: 4,
		},
		{
			gotG: []int{9, 7, 7, 9, 7, 7, 9},
			gotK: 7,
			want: 55,
		},
		{
			gotG: []int{1, 1000, 1},
			gotK: 1,
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := maxScore(testCase.gotG, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
