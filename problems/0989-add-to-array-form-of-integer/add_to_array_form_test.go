package _989_add_to_array_form_of_integer

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestAddToArrayForm(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      []int
	}{
		{
			gotNums:   []int{1, 2, 0, 0},
			gotTarget: 34,
			want:      []int{1, 2, 3, 4},
		},
		{
			gotNums:   []int{2, 7, 4},
			gotTarget: 181,
			want:      []int{4, 5, 5},
		},
		{
			gotNums:   []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			gotTarget: 1,
			want:      []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := addToArrayForm(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
