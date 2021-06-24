package _576_out_of_boundary_paths

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindPaths(t *testing.T) {
	testCases := []struct {
		gotM           int
		gotN           int
		gotMax         int
		gotStartRow    int
		gotStartColumn int
		want           int
	}{
		{
			gotM:           2,
			gotN:           2,
			gotMax:         2,
			gotStartRow:    0,
			gotStartColumn: 0,
			want:           6,
		},
		{
			gotM:           1,
			gotN:           3,
			gotMax:         3,
			gotStartRow:    0,
			gotStartColumn: 1,
			want:           12,
		},
		{
			gotM:           36,
			gotN:           5,
			gotMax:         50,
			gotStartRow:    15,
			gotStartColumn: 3,
			want:           390153306,
		},
	}

	for _, testCase := range testCases {
		actual := findPaths(testCase.gotM, testCase.gotN, testCase.gotMax, testCase.gotStartRow, testCase.gotStartColumn)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
