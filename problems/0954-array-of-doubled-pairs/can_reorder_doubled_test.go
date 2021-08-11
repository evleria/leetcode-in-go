package _954_array_of_doubled_pairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanReorderDoubled(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    bool
	}{
		{
			gotNums: []int{0, 0, 0, 0, 0, 0},
			want:    true,
		},
		{
			gotNums: []int{3, 1, 3, 6},
			want:    false,
		},
		{
			gotNums: []int{2, 1, 2, 6},
			want:    false,
		},
		{
			gotNums: []int{4, -2, 2, -4},
			want:    true,
		},
		{
			gotNums: []int{1, 2, 4, 16, 8, 4},
			want:    false,
		},
	}

	for _, testCase := range testCases {
		actual := canReorderDoubled(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.gotNums)
	}
}
