package _049_group_anagrams

import (
	"sort"
	"testing"

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

		assert.Check(t, is.DeepEqual(determine(actual), determine(testCase.want)))
	}
}

func determine(data [][]string) [][]string {
	sort.Slice(data, func(i, j int) bool {
		li, lj := len(data[i]), len(data[j])
		if li != lj {
			return li < lj
		}

		sort.Strings(data[i])
		sort.Strings(data[j])

		for k := 0; k < li; k++ {
			if data[i][k] != data[j][k] {
				return data[i][k] < data[j][k]
			}
		}

		return false
	})

	return data
}
