package _636_sort_array_by_increasing_frequency

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFrequencySort(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 1, 2, 2, 2, 3},
			want:    []int{3, 1, 1, 2, 2, 2},
		},
		{
			gotNums: []int{2, 3, 1, 3, 2},
			want:    []int{1, 3, 3, 2, 2},
		},
		{
			gotNums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			want:    []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}

	for _, testCase := range testCases {
		actual := frequencySort(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
