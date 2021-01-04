package _062_unique_paths

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUniquePaths(t *testing.T) {
	testCases := []struct {
		gotM int
		gotN int
		want int
	}{
		{
			gotM: 3,
			gotN: 7,
			want: 28,
		},
		{
			gotM: 3,
			gotN: 2,
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := uniquePaths(testCase.gotM, testCase.gotN)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
