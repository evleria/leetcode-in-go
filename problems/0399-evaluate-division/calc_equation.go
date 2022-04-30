package _399_evaluate_division

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		if _, ok := graph[equations[i][0]]; !ok {
			graph[equations[i][0]] = make(map[string]float64)
		}
		if _, ok := graph[equations[i][1]]; !ok {
			graph[equations[i][1]] = make(map[string]float64)
		}
		graph[equations[i][0]][equations[i][1]] = values[i]
		graph[equations[i][1]][equations[i][0]] = 1 / values[i]
	}

	var dfs func(start, end string, value float64, visited map[string]bool) (float64, bool)
	dfs = func(start, end string, value float64, visited map[string]bool) (float64, bool) {
		visited[start] = true
		defer delete(visited, start)

		if start == end {
			return value, true
		}

		for candidate, mul := range graph[start] {
			if visited[candidate] {
				continue
			}

			if res, ok := dfs(candidate, end, value*mul, visited); ok {
				return res, true
			}
		}

		return 0, false
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		if _, ok := graph[q[0]]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := graph[q[1]]; !ok {
			res[i] = -1
			continue
		}
		if r, ok := dfs(q[0], q[1], 1, map[string]bool{}); ok {
			res[i] = r
		} else {
			res[i] = -1
		}
	}

	return res
}
