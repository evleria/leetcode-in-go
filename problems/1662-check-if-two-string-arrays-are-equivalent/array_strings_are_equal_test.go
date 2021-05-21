package _662_check_if_two_string_arrays_are_equivalent

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestArrayStringsAreEqual(t *testing.T) {
	testCases := []struct {
		gotW1 []string
		gotW2 []string
		want  bool
	}{
		{
			gotW1: []string{"ab", "c"},
			gotW2: []string{"a", "bc"},
			want:  true,
		},
		{
			gotW1: []string{"a", "cb"},
			gotW2: []string{"ab", "c"},
			want:  false,
		},
	}

	for _, testCase := range testCases {
		actual := arrayStringsAreEqual(testCase.gotW1, testCase.gotW2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
