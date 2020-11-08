package structs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFromToSlice(t *testing.T) {
	testCases := [][]int{
		{},
		{0, 1, 4, 8},
	}

	for _, testCase := range testCases {
		list := FromSlice(testCase)
		slice := list.ToSlice()

		assert.Check(t, is.DeepEqual(slice, testCase))
	}
}

func TestListNode_NodeAt(t *testing.T) {
	testCases := []struct {
		gotList  []int
		gotIndex int
		want     int
		wantErr  string
	}{
		{
			gotList:  []int{1, 2, 3, 4},
			gotIndex: 2,
			want:     3,
		},
		{
			gotList:  []int{1, 2, 3, 4},
			gotIndex: 4,
			wantErr:  "out of bound",
		},
	}

	for _, testCase := range testCases {
		actual, err := FromSlice(testCase.gotList).NodeAt(testCase.gotIndex)

		if len(testCase.wantErr) != 0 {
			assert.Error(t, err, testCase.wantErr)
		} else {
			assert.NilError(t, err)
			assert.Check(t, is.Equal(actual.Val, testCase.want))
		}
	}
}
