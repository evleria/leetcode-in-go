package _383_ransom_note

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCanConstruct(t *testing.T) {
	testCases := []struct {
		gotNote     string
		gotMagazine string
		want        bool
	}{
		{
			gotNote:     "aa",
			gotMagazine: "ab",
			want:        false,
		},
		{
			gotNote:     "aa",
			gotMagazine: "aba",
			want:        true,
		},
	}

	for _, testCase := range testCases {
		actual := canConstruct(testCase.gotNote, testCase.gotMagazine)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
