package _171_excel_sheet_column_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "C",
			want: 3,
		},
		{
			got:  "AA",
			want: 27,
		},
		{
			got:  "AB",
			want: 28,
		},
	}

	for _, testCase := range testCases {
		actual := titleToNumber(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
