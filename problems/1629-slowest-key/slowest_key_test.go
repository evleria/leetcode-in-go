package _629_slowest_key

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSlowestKey(t *testing.T) {
	testCases := []struct {
		gotReleaseTimes []int
		gotKeysPressed  string
		want            byte
	}{
		{
			gotReleaseTimes: []int{9, 29, 49, 50},
			gotKeysPressed:  "cbcd",
			want:            'c',
		},
		{
			gotReleaseTimes: []int{12, 23, 36, 46, 62},
			gotKeysPressed:  "spuda",
			want:            'a',
		},
	}

	for _, testCase := range testCases {
		actual := slowestKey(testCase.gotReleaseTimes, testCase.gotKeysPressed)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
