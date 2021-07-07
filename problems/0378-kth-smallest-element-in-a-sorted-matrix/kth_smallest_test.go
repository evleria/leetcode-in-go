package _378_kth_smallest_element_in_a_sorted_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKthSmallest(t *testing.T) {
	testCases := []struct {
		got  [][]int
		gotK int
		want int
	}{
		{
			got:  [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
			gotK: 8,
			want: 13,
		},
		{
			got:  [][]int{{-5}},
			gotK: 1,
			want: -5,
		},
	}

	for _, testCase := range testCases {
		actual := kthSmallest(testCase.got, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
