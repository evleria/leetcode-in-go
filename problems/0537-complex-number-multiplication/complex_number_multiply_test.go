package _537_complex_number_multiplication

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestComplexNumberMultiply(t *testing.T) {
	testCases := []struct {
		got1 string
		got2 string
		want string
	}{
		{
			got1: "1+1i",
			got2: "1+1i",
			want: "0+2i",
		},
		{
			got1: "1+-1i",
			got2: "1+-1i",
			want: "0+-2i",
		},
	}

	for _, testCase := range testCases {
		actual := complexNumberMultiply(testCase.got1, testCase.got2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
