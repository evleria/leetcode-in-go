package _436_destination_city

func destCity(paths [][]string) string {
	m := make(map[string]bool, len(paths))
	for _, path := range paths {
		m[path[1]] = true
	}
	for _, path := range paths {
		delete(m, path[0])
	}
	for city := range m {
		return city
	}
	return ""
}
