package _049_group_anagrams

import (
	"testing"

	"github.com/evleria/leetcode-in-go/utils"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		got  []string
		want [][]string
	}{
		{
			got: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			got: []string{""},
			want: [][]string{
				{""},
			},
		},
		{
			got: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
	}

	for _, testCase := range testCases {
		actual := groupAnagrams(testCase.got)

		assert.Check(t, is.DeepEqual(utils.DetermineSliceOfSlicesOfStrings(actual), utils.DetermineSliceOfSlicesOfStrings(testCase.want)))
	}
}
