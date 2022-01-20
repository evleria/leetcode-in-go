package _875_koko_eating_bananas

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinEatingSpeed(t *testing.T) {
	testCases := []struct {
		got  []int
		gotH int
		want int
	}{
		{
			got:  []int{3, 6, 7, 11},
			gotH: 8,
			want: 4,
		},
		{
			got:  []int{30, 11, 23, 4, 20},
			gotH: 5,
			want: 30,
		},
		{
			got:  []int{30, 11, 23, 4, 20},
			gotH: 6,
			want: 23,
		},
	}

	for _, testCase := range testCases {
		actual := minEatingSpeed(testCase.got, testCase.gotH)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
