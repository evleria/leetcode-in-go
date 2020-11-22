package _011_conteiner_with_most_water

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxArea(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 1},
			want: 1,
		},
		{
			got:  []int{4, 3, 2, 1, 4},
			want: 16,
		},
		{
			got:  []int{1, 2, 1},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := maxArea(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
