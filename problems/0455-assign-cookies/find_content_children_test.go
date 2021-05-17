package _455_assign_cookies

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	testCases := []struct {
		gotG []int
		gotS []int
		want int
	}{
		{
			gotG: []int{1, 2, 3},
			gotS: []int{1, 1},
			want: 1,
		},
		{
			gotG: []int{1, 2},
			gotS: []int{1, 2, 3},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := findContentChildren(testCase.gotG, testCase.gotS)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
