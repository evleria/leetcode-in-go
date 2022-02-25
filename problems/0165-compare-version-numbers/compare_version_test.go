package _165_compare_version_numbers

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCompareVersion(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want int
	}{
		{
			got1: "1.01",
			got2: "1.001",
			want: 0,
		},
		{
			got1: "1.0",
			got2: "1.0.0",
			want: 0,
		},
		{
			got1: "0.1",
			got2: "1.1",
			want: -1,
		},
		{
			got1: "1.0",
			got2: "1.1",
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := compareVersion(testCase.got1, testCase.got2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
