package _551_minimum_operations_to_make_array_equal

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinOperations(t *testing.T) {
	testCases := []struct {
		got  int
		want int
	}{
		{
			got:  3,
			want: 2,
		},
		{
			got:  6,
			want: 9,
		},
	}

	for _, testCase := range testCases {
		actual := minOperations(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
