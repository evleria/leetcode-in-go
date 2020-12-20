package _190_reverse_bits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseBits(t *testing.T) {
	testCases := []struct {
		got  uint32
		want uint32
	}{
		{
			got:  0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
		{
			got:  0b11111111111111111111111111111101,
			want: 0b10111111111111111111111111111111,
		},
	}

	for _, testCase := range testCases {
		actual := reverseBits(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
