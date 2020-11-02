package _007_reverse_integer

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverse(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  321,
			want: 123,
		},
		{
			got:  120,
			want: 21,
		},
		{
			got:  0,
			want: 0,
		},
		{
			got:  -123,
			want: -321,
		},
		{
			got:  1534236469,
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := reverse(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
