package _387_first_unique_character_in_a_string

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFirstUniqueCharacter(t *testing.T) {
	testCases := []struct {
		got  string
		want int
	}{
		{
			got:  "leetcode",
			want: 0,
		},
		{
			got:  "loveleetcode",
			want: 2,
		},
		{
			got:  "aabb",
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := firstUniqChar(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
