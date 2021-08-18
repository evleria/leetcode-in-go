package _091_decode_ways

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumDecodings(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "505",
			want: 0,
		},
		{
			got:  "12",
			want: 2,
		},
		{
			got:  "226",
			want: 3,
		},
		{
			got:  "0",
			want: 0,
		},
		{
			got:  "06",
			want: 0,
		},
		{
			got:  "1110",
			want: 2,
		},
		{
			got:  "11100",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := numDecodings(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
