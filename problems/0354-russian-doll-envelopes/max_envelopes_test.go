package _354_russian_doll_envelopes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxEnvelopes(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{1, 2}, {2, 3}, {3, 4}, {3, 5}, {4, 5}, {5, 5}, {5, 6}, {6, 7}, {7, 8}},
			want: 7,
		},
		{
			got:  [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
			want: 3,
		},
		{
			got:  [][]int{{1, 1}, {1, 1}, {1, 1}},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := maxEnvelopes(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
