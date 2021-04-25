package _914_x_of_a_kind_in_a_deck_of_cards

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestHasGroupsSize(t *testing.T) {
	testCases := []struct {
		got  []int
		want bool
	}{
		{
			got:  []int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
			want: true,
		},
		{
			got:  []int{1, 2, 3, 4, 4, 3, 2, 1},
			want: true,
		},
		{
			got:  []int{1, 1, 1, 2, 2, 2, 3, 3},
			want: false,
		},
		{
			got:  []int{1},
			want: false,
		},
	}

	for _, testCase := range testCases {
		actual := hasGroupsSizeX(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
