package _399_evaluate_division

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCalcEquation(t *testing.T) {
	testCases := []struct {
		gotEquations [][]string
		gotValues    []float64
		gotQueries   [][]string
		want         []float64
	}{
		{
			gotEquations: [][]string{
				{"a", "b"},
				{"b", "c"},
			},
			gotValues: []float64{2.0, 3.0},
			gotQueries: [][]string{
				{"a", "c"},
				{"b", "a"},
				{"a", "e"},
				{"a", "a"},
				{"x", "x"},
			},
			want: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
	}

	for _, testCase := range testCases {
		actual := calcEquation(testCase.gotEquations, testCase.gotValues, testCase.gotQueries)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
