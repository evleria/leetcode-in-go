package _486_xor_operation_in_an_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestXorOperation(t *testing.T) {
	testCases := []struct {
		gotN     int
		gotStart int
		want     int
	}{
		{
			gotN:     5,
			gotStart: 0,
			want:     8,
		},
		{
			gotN:     4,
			gotStart: 3,
			want:     8,
		},
	}

	for _, testCase := range testCases {
		actual := xorOperation(testCase.gotN, testCase.gotStart)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
