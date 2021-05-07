package _583_delete_operation_for_two_strings

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinDistance(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want int
	}{
		{
			got1: "leetcode",
			got2: "xetcdf",
			want: 6,
		},

		{
			got1: "sea",
			got2: "eat",
			want: 2,
		},
		{
			got1: "leetcode",
			got2: "etco",
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := minDistance(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
