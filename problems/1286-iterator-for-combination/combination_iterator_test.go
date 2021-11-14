package _286_iterator_for_combination

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCombinationIterator(t *testing.T) {
	testCases := []struct {
		gotCharacters string
		gotLength     int
		want          []string
	}{
		{
			gotCharacters: "abc",
			gotLength:     2,
			want:          []string{"ab", "ac", "bc"},
		},
	}

	for _, testCase := range testCases {
		iterator := Constructor(testCase.gotCharacters, testCase.gotLength)

		actual := []string{iterator.Next()}
		for iterator.HasNext() {
			actual = append(actual, iterator.Next())
		}

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
