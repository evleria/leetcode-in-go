package _869_reordered_power_of_2

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRecorderPowerOf2(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    bool
	}{
		{
			gotNums: 1,
			want:    true,
		},
		{
			gotNums: 10,
			want:    false,
		},
		{
			gotNums: 46,
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := reorderedPowerOf2(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
