package _609_find_duplicate_file_in_system

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	testCases := []struct {
		got  []string
		want [][]string
	}{
		{
			got: []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			want: [][]string{
				{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
				{"root/a/1.txt", "root/c/3.txt"},
			},
		},
		{
			got: []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"},
			want: [][]string{
				{"root/a/2.txt", "root/c/d/4.txt"},
				{"root/a/1.txt", "root/c/3.txt"},
			},
		},
	}

	for _, testCase := range testCases {
		actual := findDuplicate(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
