package _946_validate_stack_sequences

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestValidateStackSequences(t *testing.T) {
	testCases := []struct {
		gotPushed []int
		gotPopped []int
		want      bool
	}{
		{
			gotPushed: []int{1, 2, 3, 4, 5},
			gotPopped: []int{4, 5, 3, 2, 1},
			want:      true,
		},
		{
			gotPushed: []int{1, 2, 3, 4, 5},
			gotPopped: []int{4, 3, 5, 1, 2},
			want:      false,
		},
	}

	for _, testCase := range testCases {
		actual := validateStackSequences(testCase.gotPushed, testCase.gotPopped)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
