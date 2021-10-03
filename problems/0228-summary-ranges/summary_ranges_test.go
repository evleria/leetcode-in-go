package _228_summary_ranges

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSummaryRanges(t *testing.T) {
	testCases := []struct {
		got  []int
		want []string
	}{
		{
			got:  []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			got:  []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			got:  []int{},
			want: []string{},
		},
		{
			got:  []int{-1},
			want: []string{"-1"},
		},
		{
			got:  []int{0},
			want: []string{"0"},
		},
	}

	for _, testCase := range testCases {
		actual := summaryRanges(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
