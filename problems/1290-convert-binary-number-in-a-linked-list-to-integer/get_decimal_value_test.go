package _290_convert_binary_number_in_a_linked_list_to_integer

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetDecimalValue(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{1, 0, 1},
			want: 5,
		},
		{
			got:  []int{0},
			want: 0,
		},
		{
			got:  []int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			want: 18880,
		},
	}

	for _, testCase := range testCases {
		actual := getDecimalValue(LinkedListFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
