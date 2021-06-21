package _313_decompress_run_length_encoded_list

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDecompressRleList(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 2, 3, 4},
			want:    []int{2, 4, 4, 4},
		},
		{
			gotNums: []int{1, 1, 2, 3},
			want:    []int{1, 3, 3},
		},
	}

	for _, testCase := range testCases {
		actual := decompressRLElist(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
