package _331_verify_preorder_serialization_of_a_binary_tree

import "strings"

func isValidSerialization(preorder string) bool {
	diff := 1
	for _, token := range strings.Split(preorder, ",") {
		diff--
		if diff < 0 {
			return false
		}

		if token != "#" {
			diff += 2
		}
	}
	return diff == 0
}
