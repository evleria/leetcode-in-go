package _268_search_suggestions_system

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSugestedProducts(t *testing.T) {
	testCases := []struct {
		gotProduct    []string
		gotSearchWord string
		want          [][]string
	}{
		{
			gotProduct:    []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			gotSearchWord: "mouse",
			want: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
		{
			gotProduct:    []string{"havana"},
			gotSearchWord: "havana",
			want:          [][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			gotProduct:    []string{"bags", "baggage", "banner", "box", "cloths"},
			gotSearchWord: "bags",
			want: [][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"},
				{"bags"},
			},
		},
		{
			gotProduct:    []string{"havana"},
			gotSearchWord: "tatiana",
			want:          [][]string{{}, {}, {}, {}, {}, {}, {}},
		},
	}

	for _, testCase := range testCases {
		actual := suggestedProducts(testCase.gotProduct, testCase.gotSearchWord)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
