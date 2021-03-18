package _478_generate_random_point_in_a_circle

import (
	"math"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSolution(t *testing.T) {
	x, y, r := 5.5, -3.6, 8.1
	solution := Constructor(r, x, y)

	belongs := func(x0, y0 float64) bool {
		xdiff, ydiff := x-x0, y-y0
		return math.Sqrt(xdiff*xdiff+ydiff*ydiff) <= r
	}
	for i := 0; i < 100; i++ {
		point := solution.RandPoint()

		assert.Check(t, is.Equal(true, belongs(point[0], point[1])))
	}
}
