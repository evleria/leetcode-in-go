package _797_all_paths_from_source_to_target

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	paths := [][]int{}
	var dfs func(int, []int)
	dfs = func(from int, path []int) {
		if from == n-1 {
			cp := make([]int, len(path))
			copy(cp, path)
			paths = append(paths, cp)
			return
		}

		for _, dest := range graph[from] {
			dfs(dest, append(path, dest))
		}
	}

	dfs(0, append(make([]int, 0, n), 0)) // important to allocate a buffer with enough space

	return paths
}
