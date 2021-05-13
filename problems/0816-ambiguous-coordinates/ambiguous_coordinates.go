package _816_ambiguous_coordinates

import "fmt"

func ambiguousCoordinates(s string) []string {
	var result []string

	for i := 2; i <= len(s)-2; i++ {
		left, right := s[1:i], s[i:len(s)-1]

		if isLegalPart(left) && isLegalPart(right) {
			leftPartitions, rightPartitions := getAllPartitions(left), getAllPartitions(right)

			for _, leftPartition := range leftPartitions {
				for _, rightPartition := range rightPartitions {
					result = append(result, fmt.Sprintf("(%s, %s)", leftPartition, rightPartition))
				}
			}
		}
	}

	return result
}

func isLegalPart(s string) bool {
	return len(s) == 1 || s[0] != '0' || s[len(s)-1] != '0'
}

func getAllPartitions(s string) []string {
	switch {
	case len(s) == 1 || s[len(s)-1] == '0':
		return []string{s}
	case s[0] == '0':
		return []string{fmt.Sprintf("0.%s", s[1:])}
	default:
		result := make([]string, 1, len(s))
		result[0] = s
		for i := 1; i < len(s); i++ {
			result = append(result, fmt.Sprintf("%s.%s", s[:i], s[i:]))
		}
		return result
	}
}
