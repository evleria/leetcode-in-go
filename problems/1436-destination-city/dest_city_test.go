package _436_destination_city

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDestCity(t *testing.T) {
	testCases := []struct {
		got  [][]string
		want string
	}{
		{
			got:  [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
			want: "Sao Paulo",
		},
		{
			got:  [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
			want: "A",
		},
	}

	for _, testCase := range testCases {
		actual := destCity(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
