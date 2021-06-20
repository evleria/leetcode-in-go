package _431_kids_with_the_greatest_number_of_candies

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		gotCandies      []int
		gotExtraCandies int
		want            []bool
	}{
		{
			gotCandies:      []int{2, 3, 5, 1, 3},
			gotExtraCandies: 3,
			want:            []bool{true, true, true, false, true},
		},
		{
			gotCandies:      []int{4, 2, 1, 1, 2},
			gotExtraCandies: 1,
			want:            []bool{true, false, false, false, false},
		},
		{
			gotCandies:      []int{12, 1, 12},
			gotExtraCandies: 10,
			want:            []bool{true, false, true},
		},
	}

	for _, testCase := range testCases {
		actual := kidsWithCandies(testCase.gotCandies, testCase.gotExtraCandies)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
