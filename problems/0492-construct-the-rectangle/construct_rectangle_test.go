package _492_construct_the_rectangle

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	testCases := []struct {
		got  int
		want []int
	}{
		{
			got:  4,
			want: []int{2, 2},
		},
		{
			got:  37,
			want: []int{37, 1},
		},
		{
			got:  122122,
			want: []int{427, 286},
		},
	}

	for _, testCase := range testCases {
		actual := constructRectangle(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
