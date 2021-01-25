package _268_missing_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMissingNumber(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{3, 0, 1},
			want:    2,
		},
		{
			gotNums: []int{0, 1},
			want:    2,
		},
		{
			gotNums: []int{0},
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := missingNumber(testCase.gotNums)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
