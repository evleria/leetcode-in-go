package _775_global_and_local_inversions

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsIdealPermutation(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1, 0, 2},
			want: true,
		},
		{
			got:  []int{1, 2, 0},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := isIdealPermutation(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
