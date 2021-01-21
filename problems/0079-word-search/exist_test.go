package _079_word_search

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestExist(t *testing.T) {
	testCases := []struct {
		gotBoard [][]byte
		gotWord  string
		want     bool
	}{
		{
			gotBoard: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			gotWord: "ABCCED",
			want:    true,
		},
		{
			gotBoard: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			gotWord: "SEE",
			want:    true,
		},
		{
			gotBoard: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			gotWord: "ABCCCD",
			want:    false,
		},
		{
			gotBoard: [][]byte{
				{'C', 'A', 'A'},
				{'A', 'A', 'A'},
				{'B', 'C', 'D'},
			},
			gotWord: "AAB",
			want:    true,
		},
	}

	for _, testCase := range testCases {
		actual := exist(testCase.gotBoard, testCase.gotWord)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
