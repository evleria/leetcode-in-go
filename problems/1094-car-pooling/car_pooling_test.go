package _094_car_pooling

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCarPooling(t *testing.T) {
	testCases := []struct {
		gotT [][]int
		gotC int
		want bool
	}{
		{
			gotT: [][]int{{2, 1, 5}, {3, 3, 7}},
			gotC: 4,
			want: false,
		},
		{
			gotT: [][]int{{2, 1, 5}, {3, 3, 7}},
			gotC: 5,
			want: true,
		},
	}

	for _, testCase := range testCases {
		actual := carPooling(testCase.gotT, testCase.gotC)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
