package _137_n_th_tribonacci_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTribonacci(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  4,
			want: 4,
		},
		{
			got:  25,
			want: 1389537,
		},
	}

	for _, testCase := range testCases {
		actual := tribonacci(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
