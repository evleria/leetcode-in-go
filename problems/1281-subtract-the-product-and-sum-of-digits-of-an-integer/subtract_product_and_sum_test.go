package _281_subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSubtractProductAndSum(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  234,
			want: 15,
		},
		{
			got:  4421,
			want: 21,
		},
	}

	for _, testCase := range testCases {
		actual := subtractProductAndSum(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
