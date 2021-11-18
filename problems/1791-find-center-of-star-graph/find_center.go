package _791_find_center_of_star_graph

func findCenter(edges [][]int) int {
	e1, e2 := edges[0], edges[1]
	if e1[0] == e2[0] || e1[0] == e2[1] {
		return e1[0]
	}
	return e1[1]
}
