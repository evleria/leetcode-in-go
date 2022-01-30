package _189_rotate_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRotate(t *testing.T) {
	testCases := []struct {
		got  []int
		gotK int
		want []int
	}{
		{
			got:  []int{1, 2, 3, 4, 5, 6, 7},
			gotK: 3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			got:  []int{-1, -100, 3, 99},
			gotK: 2,
			want: []int{3, 99, -1, -100},
		},
		{
			got:  []int{1, 2, 3, 4, 5, 6},
			gotK: 4,
			want: []int{3, 4, 5, 6, 1, 2},
		},
		{
			got:  []int{1},
			gotK: 0,
			want: []int{1},
		},
		{
			got:  []int{1},
			gotK: 1,
			want: []int{1},
		},
	}

	for _, testCase := range testCases {
		rotate(testCase.got, testCase.gotK)
		assert.Check(t, is.DeepEqual(testCase.got, testCase.want))
	}
}
