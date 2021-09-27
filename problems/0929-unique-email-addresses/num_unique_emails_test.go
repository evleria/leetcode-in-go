package _929_unique_email_addresses

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumUniqueEmails(t *testing.T) {
	testCases := []struct {
		got  []string
		want int
	}{
		{
			got: []string{
				"test.email+alex@leetcode.com",
				"test.e.mail+bob.cathy@leetcode.com",
				"testemail+david@lee.tcode.com"},
			want: 2,
		},
		{
			got:  []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"},
			want: 3,
		},
	}

	for _, testCase := range testCases {
		actual := numUniqueEmails(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
