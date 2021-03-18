package _478_generate_random_point_in_a_circle

import (
	"math"
	"math/rand"
)

type Solution struct {
	x, y, r float64
}

func Constructor(radius, x, y float64) Solution {
	return Solution{x, y, radius}
}

func (s *Solution) RandPoint() []float64 {
	a, d := rand.Float64()*2*math.Pi, math.Sqrt(rand.Float64())*s.r
	return []float64{s.x + d*math.Cos(a), s.y + d*math.Sin(a)}
}
