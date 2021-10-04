package _323_maximum_69_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximum69Number(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  9669,
			want: 9969,
		},
		{
			got:  9996,
			want: 9999,
		},
		{
			got:  9999,
			want: 9999,
		},
	}

	for _, testCase := range testCases {
		actual := maximum69Number(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
