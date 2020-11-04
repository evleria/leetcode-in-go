package _066_plus_one

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, 2, 3},
			want: []int{1, 2, 4},
		},
		{
			got:  []int{1, 2, 9},
			want: []int{1, 3, 0},
		},
		{
			got:  []int{1, 9, 9, 9},
			want: []int{2, 0, 0, 0},
		},
		{
			got:  []int{9, 9, 9},
			want: []int{1, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := plusOne(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
