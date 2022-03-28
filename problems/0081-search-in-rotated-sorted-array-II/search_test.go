package _081_search_in_rotated_sorted_array_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		got       []int
		gotTarget int
		want      bool
	}{
		{
			got:       []int{2, 5, 6, 0, 0, 1, 2},
			gotTarget: 0,
			want:      true,
		},
		{
			got:       []int{2, 5, 6, 0, 0, 1, 2},
			gotTarget: 3,
			want:      false,
		},
	}

	for _, testCase := range testCases {
		actual := search(testCase.got, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
