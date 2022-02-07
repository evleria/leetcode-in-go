package _154_keep_multiplying_found_values_by_two

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindFinalValue(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{5, 3, 6, 1, 12},
			gotTarget: 3,
			want:      24,
		},
		{
			gotNums:   []int{2, 7, 9},
			gotTarget: 4,
			want:      4,
		},
	}

	for _, testCase := range testCases {
		actual := findFinalValue(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
