package _009_complement_of_base_10_integer

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBitwiseComplement(t *testing.T) {
	testCases := []struct {
		gotN int
		want int
	}{
		{
			gotN: 5,
			want: 2,
		},
		{
			gotN: 7,
			want: 0,
		},
		{
			gotN: 10,
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := bitwiseComplement(testCase.gotN)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
