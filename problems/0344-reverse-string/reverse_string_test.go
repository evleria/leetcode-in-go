package _344_reverse_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		got  []byte
		want []byte
	}{
		{
			got:  []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
	}

	for _, testCase := range testCases {
		reverseString(testCase.got)

		assert.Check(t, is.DeepEqual(testCase.got, testCase.want))
	}
}
