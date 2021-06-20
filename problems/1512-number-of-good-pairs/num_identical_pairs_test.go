package _512_number_of_good_pairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumIdenticalPairs(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3, 1, 1, 3},
			want:    4,
		},
		{
			gotNums: []int{1, 1, 1, 1},

			want: 6,
		},
		{
			gotNums: []int{1, 2, 3},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := numIdenticalPairs(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
