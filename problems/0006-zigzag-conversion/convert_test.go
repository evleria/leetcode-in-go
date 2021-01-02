package _006_zigzag_conversion

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		gotString  string
		gotNumRows int
		want       string
	}{
		{
			gotString:  "PAYPALISHIRING",
			gotNumRows: 3,
			want:       "PAHNAPLSIIGYIR",
		},
		{
			gotString:  "PAYPALISHIRING",
			gotNumRows: 4,
			want:       "PINALSIGYAHRPI",
		},
		{
			gotString:  "A",
			gotNumRows: 1,
			want:       "A",
		},
	}

	for _, testCase := range testCases {
		actual := convert(testCase.gotString, testCase.gotNumRows)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
