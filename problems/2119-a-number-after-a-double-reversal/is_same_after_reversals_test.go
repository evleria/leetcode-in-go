package _119_a_number_after_a_double_reversal

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestIsSameAfterReversals(t *testing.T) {
	testCases := []struct {
		got  int
		want bool
	}{
		{
			got:  526,
			want: true,
		},
		{
			got:  1800,
			want: false,
		},
		{
			got:  0,
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := isSameAfterReversals(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
