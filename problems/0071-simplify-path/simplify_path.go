package _071_simplify_path

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0)
	splits := strings.Split(path, "/")
	for i := 0; i < len(splits); i++ {
		switch splits[i] {
		case "..":
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		case "", ".":
			continue
		default:
			stack = append(stack, splits[i])
		}
	}
	result := strings.Join(stack, "/")
	return "/" + result
}
