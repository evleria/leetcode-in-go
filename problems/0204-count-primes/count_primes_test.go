package _204_count_primes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountPrimes(t *testing.T) {
	testCases := []struct {
		gotN int
		want int
	}{
		{
			gotN: 10,
			want: 4,
		},
		{
			gotN: 0,
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := countPrimes(testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
