package _629_k_inverse_pairs_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKInversePairs(t *testing.T) {
	testCases := []struct {
		gotN int
		gotK int
		want int
	}{
		{
			gotN: 3,
			gotK: 0,
			want: 1,
		},
		{
			gotN: 3,
			gotK: 1,
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := kInversePairs(testCase.gotN, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
