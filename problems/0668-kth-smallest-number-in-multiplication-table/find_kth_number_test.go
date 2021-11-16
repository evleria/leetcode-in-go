package _668_kth_smallest_number_in_multiplication_table

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindKthNumber(t *testing.T) {
	testCases := []struct {
		gotM int
		gotN int
		gotK int
		want int
	}{
		{
			gotM: 3,
			gotN: 3,
			gotK: 5,
			want: 3,
		},
		{
			gotM: 2,
			gotN: 3,
			gotK: 6,
			want: 6,
		},
	}

	for _, testCase := range testCases {
		actual := findKthNumber(testCase.gotM, testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
