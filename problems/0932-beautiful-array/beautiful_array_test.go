package _932_beautiful_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBeautifulArray(t *testing.T) {
	testCases := []struct {
		got  int
		want []int
	}{
		{
			got:  4,
			want: []int{1, 3, 2, 4},
		},
		{
			got:  5,
			want: []int{1, 5, 3, 2, 4},
		},
	}

	for _, testCase := range testCases {
		actual := beautifulArray(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
