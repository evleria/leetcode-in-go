package _941_valid_mountain_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestValidMountainArray(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{2, 1},
			want: false,
		},
		{
			got:  []int{3, 5, 5},
			want: false,
		},
		{
			got:  []int{0, 3, 2, 1},
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := validMountainArray(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
