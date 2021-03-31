package _936_stamping_the_sequence

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMovesToStamp(t *testing.T) {
	testCases := []struct {
		gotStamp  string
		gotTarget string
		want      []int
	}{
		{
			gotStamp:  "abc",
			gotTarget: "ababc",
			want:      []int{0, 2},
		},
		{
			gotStamp:  "abca",
			gotTarget: "aabcaca",
			want:      []int{3, 0, 1},
		},
	}

	for _, testCase := range testCases {
		actual := movesToStamp(testCase.gotStamp, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
