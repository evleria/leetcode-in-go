package _347_top_k_frequent_elements

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    []int
	}{
		{
			gotNums: []int{1, 1, 1, 2, 2, 3},
			gotK:    2,
			want:    []int{1, 2},
		},
		{
			gotNums: []int{1},
			gotK:    1,
			want:    []int{1},
		},
	}

	for _, testCase := range testCases {
		actual := topKFrequent(testCase.gotNums, testCase.gotK)
		sort.Ints(actual)
		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotNums)
	}
}
