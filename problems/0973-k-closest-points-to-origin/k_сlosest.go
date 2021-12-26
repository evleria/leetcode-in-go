package _973_k_closest_points_to_origin

import (
	"math"
	"sort"
)

type wrapper struct {
	points    [][]int
	distances []float64
}

func (w *wrapper) Len() int {
	return len(w.points)
}

func (w *wrapper) Less(i, j int) bool {
	return w.distances[i] < w.distances[j]
}

func (w *wrapper) Swap(i, j int) {
	w.points[i], w.points[j] = w.points[j], w.points[i]
	w.distances[i], w.distances[j] = w.distances[j], w.distances[i]
}

func kClosest(points [][]int, k int) [][]int {
	distances := make([]float64, len(points))
	for i, p := range points {
		distances[i] = distanceToOrigin(p)
	}

	w := wrapper{points, distances}
	sort.Sort(&w)

	return points[:k]
}

func distanceToOrigin(point []int) float64 {
	x, y := point[0], point[1]
	return math.Sqrt(float64(x*x + y*y))
}
