package _763_partition_labels

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPartitionLabels(t *testing.T) {
	testCases := []struct {
		gotNums string
		want    []int
	}{
		{
			gotNums: "ababcbacadefegdehijhklij",
			want:    []int{9, 7, 8},
		},
		{
			gotNums: "eccbbbbdec",
			want:    []int{10},
		},
	}

	for _, testCase := range testCases {
		actual := partitionLabels(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
