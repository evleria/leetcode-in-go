package _476_number_complement

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindComplement(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    int
	}{
		{
			gotNums: 5,
			want:    2,
		},
		{
			gotNums: 1,
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := findComplement(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
