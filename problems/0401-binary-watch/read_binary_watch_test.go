package _401_binary_watch

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestReadBinaryWatch(t *testing.T) {
	testCases := []struct {
		got  int
		want []string
	}{
		{
			got:  1,
			want: []string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			got:  9,
			want: []string{},
		},
	}

	for _, testCase := range testCases {
		actual := readBinaryWatch(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
