package _837_sum_of_digits_in_base_k

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumBase(t *testing.T) {
	testCases := []struct {
		gotN int
		gotK int
		want int
	}{
		{
			gotN: 34,
			gotK: 6,
			want: 9,
		},
		{
			gotN: 10,
			gotK: 10,
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := sumBase(testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
