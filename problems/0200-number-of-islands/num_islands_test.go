package _200_number_of_islands

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumIslands(t *testing.T) {
	testCases := []struct {
		got  [][]byte
		want int
	}{
		{
			got: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			want: 1,
		},
		{
			got: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := numIslands(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
