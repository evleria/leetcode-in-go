package _902_numbers_at_most_n_given_digit_set

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAtMostNGivenDigitSet(t *testing.T) {
	testCases := []struct {
		got  []string
		gotN int
		want int
	}{
		{
			got:  []string{"1", "3", "5", "7"},
			gotN: 100,
			want: 20,
		},
		{
			got:  []string{"1", "4", "9"},
			gotN: 1000000000,
			want: 29523,
		},
	}

	for _, testCase := range testCases {
		actual := atMostNGivenDigitSet(testCase.got, testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
