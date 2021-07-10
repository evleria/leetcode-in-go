package _639_decode_ways_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumLengthEncoding(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "1",
			want: 1,
		},
		{
			got:  "34",
			want: 1,
		},
		{
			got:  "134",
			want: 2,
		},
		{
			got:  "12",
			want: 2,
		},
		{
			got:  "10",
			want: 1,
		},
		{
			got:  "21",
			want: 2,
		},
		{
			got:  "28",
			want: 1,
		},
		{
			got:  "20",
			want: 1,
		},
		{
			got:  "111",
			want: 3,
		},
		{
			got:  "*",
			want: 9,
		}, /*
			{
				got:  "1*",
				want: 18,
			},
			{
				got:  "2*",
				want: 15,
			},*/
	}

	for _, testCase := range testCases {
		actual := numDecodings(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
