package _180_count_integers_with_even_digit_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountEven(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  4,
			want: 2,
		},
		{
			got:  30,
			want: 14,
		},
	}

	for _, testCase := range testCases {
		actual := countEven(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
