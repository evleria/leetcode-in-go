package _504_base_7

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConvertToBase7(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    string
	}{
		{
			gotNums: 0,
			want:    "0",
		},
		{
			gotNums: 100,
			want:    "202",
		},
		{
			gotNums: -7,
			want:    "-10",
		},
	}

	for _, testCase := range testCases {
		actual := convertToBase7(testCase.gotNums)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
