package _169_count_operations_to_obtain_zero

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountOperations(t *testing.T) {
	testCases := []struct {
		gotNum1 int
		gotNum2 int
		want    int
	}{
		{
			gotNum1: 2,
			gotNum2: 3,
			want:    3,
		},
		{
			gotNum1: 10,
			gotNum2: 10,
			want:    1,
		},
	}

	for _, testCase := range testCases {
		actual := countOperations(testCase.gotNum1, testCase.gotNum2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
