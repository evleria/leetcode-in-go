package _178_number_of_valid_words_for_each_puzzle

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindNumOfValidWords(t *testing.T) {
	testCases := []struct {
		gotWords   []string
		gotPuzzles []string
		want       []int
	}{
		{
			gotWords:   []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"},
			gotPuzzles: []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"},
			want:       []int{1, 1, 3, 2, 4, 0},
		},
		{
			gotWords:   []string{"apple", "pleas", "please"},
			gotPuzzles: []string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"},
			want:       []int{0, 1, 3, 2, 0},
		},
	}

	for _, testCase := range testCases {
		actual := findNumOfValidWords(testCase.gotWords, testCase.gotPuzzles)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
