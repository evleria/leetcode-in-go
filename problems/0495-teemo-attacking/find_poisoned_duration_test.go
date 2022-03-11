package _495_teemo_attacking

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindPoisonedDuration(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{1, 4},
			gotTarget: 2,
			want:      4,
		},
		{
			gotNums:   []int{1, 2},
			gotTarget: 2,
			want:      3,
		},
	}

	for _, testCase := range testCases {
		actual := findPoisonedDuration(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
