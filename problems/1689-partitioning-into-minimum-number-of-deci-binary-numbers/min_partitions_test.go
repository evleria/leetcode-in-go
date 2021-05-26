package _689_partitioning_into_minimum_number_of_deci_binary_numbers

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinPartition(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "32",
			want: 3,
		},
		{
			got:  "82734",
			want: 8,
		},
		{
			got:  "27346209830709182346",
			want: 9,
		},
	}

	for _, testCase := range testCases {
		actual := minPartitions(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
