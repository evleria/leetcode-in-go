package _695_maximum_erasure_value

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaximumUniqueSubbarray(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{4, 2, 4, 5, 6},
			want: 17,
		},
		{
			got:  []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			want: 8,
		},
	}

	for _, testCase := range testCases {
		actual := maximumUniqueSubarray(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
