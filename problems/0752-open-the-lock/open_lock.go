package _752_open_the_lock

import "strconv"

func openLock(deadends []string, target string) int {
	t, _ := strconv.Atoi(target)
	if t == 0 {
		return 0
	}

	deadendsMap := [10_000]bool{}
	for _, deadend := range deadends {
		n, _ := strconv.Atoi(deadend)
		deadendsMap[n] = true
	}
	if deadendsMap[0] {
		return -1
	}

	counter, queue, visited := 0, []int{0}, [10_000]bool{}
	visited[0] = true
	for len(queue) != 0 {
		counter++
		for _, combination := range queue {
			queue = queue[1:]
			for _, candidate := range getAllCandidates(combination) {
				if candidate == t {
					return counter
				}

				if !deadendsMap[candidate] && !visited[candidate] {
					visited[candidate] = true
					queue = append(queue, candidate)
				}
			}
		}
	}

	return -1
}

func getAllCandidates(combination int) []int {
	result := make([]int, 0, 8)
	for place := 1; place <= 1_000; place *= 10 {
		d := (combination / place) % 10
		switch d {
		case 0:
			result = append(result, combination+9*place, combination+place)
		case 9:
			result = append(result, combination-9*place, combination-place)
		default:
			result = append(result, combination-place, combination+place)
		}
	}
	return result
}
