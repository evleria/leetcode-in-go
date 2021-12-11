package _878_nth_magical_number

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestNthMagicalNumber(t *testing.T) {
	testCases := []struct {
		gotN int
		gotA int
		gotB int
		want int
	}{
		{
			gotN: 1,
			gotA: 2,
			gotB: 3,
			want: 2,
		},
		{
			gotN: 4,
			gotA: 2,
			gotB: 3,
			want: 6,
		},
		{
			gotN: 5,
			gotA: 2,
			gotB: 4,
			want: 10,
		},
	}

	for _, testCase := range testCases {
		actual := nthMagicalNumber(testCase.gotN, testCase.gotA, testCase.gotB)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
