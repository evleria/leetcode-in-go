package _108_defanging_an_ip_address

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestDefangIPaddr(t *testing.T) {
	testCases := []struct {
		got  string
		want string
	}{
		{
			got:  "1.1.1.1",
			want: "1[.]1[.]1[.]1",
		},
		{
			got:  "255.100.50.0",
			want: "255[.]100[.]50[.]0",
		},
	}

	for _, testCase := range testCases {
		actual := defangIPaddr(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
