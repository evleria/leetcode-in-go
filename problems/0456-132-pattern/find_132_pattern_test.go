package _456_132_pattern

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFind132Pattern(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1, 2, 3, 4},
			want: false,
		},
		{
			got:  []int{3, 1, 4, 2},
			want: true,
		},
		{
			got:  []int{-1, 3, 2, 0},
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := find132pattern(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
