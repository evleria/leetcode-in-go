package _169_majority_element

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMajorityElement(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1},
			want: 1,
		},
		{
			got:  []int{2, 3, 3},
			want: 3,
		},
		{
			got:  []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := majorityElement(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
