package _212_word_search_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindWords(t *testing.T) {
	testCases := []struct {
		gotBoard [][]byte
		gotWord  []string
		want     []string
	}{
		{
			gotBoard: [][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			gotWord: []string{"oath", "pea", "eat", "rain"},
			want:    []string{"oath", "eat"},
		},
		{
			gotBoard: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			gotWord: []string{"abcb", "aba"},
			want:    []string{},
		},
	}

	for _, testCase := range testCases {
		actual := findWords(testCase.gotBoard, testCase.gotWord)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
