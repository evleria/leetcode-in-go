package _010_pairs_of_songs_with_total_durations_divisible_by_60

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumPairsDivisibleBy60(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{30, 20, 150, 100, 40},
			want: 3,
		},
		{
			got:  []int{60, 60, 60},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := numPairsDivisibleBy60(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
