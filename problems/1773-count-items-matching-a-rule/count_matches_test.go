package _773_count_items_matching_a_rule

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountMatches(t *testing.T) {
	testCases := []struct {
		gotItems [][]string
		gotKey   string
		gotValue string
		want     int
	}{
		{
			gotItems: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			gotKey:   "color",
			gotValue: "silver",
			want:     1,
		},
		{
			gotItems: [][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "phone"},
				{"phone", "gold", "iphone"},
			},
			gotKey:   "type",
			gotValue: "phone",
			want:     2,
		},
	}

	for _, testCase := range testCases {
		actual := countMatches(testCase.gotItems, testCase.gotKey, testCase.gotValue)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
