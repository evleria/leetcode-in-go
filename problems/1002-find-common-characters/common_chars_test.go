package _002_find_common_characters

import (
	"sort"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCommonChars(t *testing.T) {
	testCases := []struct {
		got  []string
		want []string
	}{
		{
			got:  []string{"bella", "label", "roller"},
			want: []string{"e", "l", "l"},
		},
		{
			got:  []string{"cool", "lock", "cook"},
			want: []string{"c", "o"},
		},
	}

	for _, testCase := range testCases {
		actual := commonChars(testCase.got)
		sort.Strings(actual)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
