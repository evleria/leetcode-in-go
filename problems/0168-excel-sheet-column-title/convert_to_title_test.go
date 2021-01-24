package _168_excel_sheet_column_title

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConvertToTitle(t *testing.T) {
	testCases := []struct {
		got  int
		want string
	}{
		{
			got:  26,
			want: "Z",
		},
		{
			got:  28,
			want: "AB",
		},
		{
			got:  701,
			want: "ZY",
		},
	}

	for _, testCase := range testCases {
		actual := convertToTitle(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
