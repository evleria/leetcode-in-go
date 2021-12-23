package _210_course_schedule_II

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges, indegree := make([][]int, numCourses), make([]int, numCourses)
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indegree[v[0]]++
	}
	queue, res := []int{}, []int{}
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		numCourses--
		res = append(res, cur)
		for _, v := range edges[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	if numCourses != 0 {
		return []int{}
	}
	return res
}
