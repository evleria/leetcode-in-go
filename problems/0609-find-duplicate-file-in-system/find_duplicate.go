package _609_find_duplicate_file_in_system

import (
	"strings"
)

func findDuplicate(paths []string) [][]string {
	groups := make(map[string][]string)
	for _, path := range paths {
		i := strings.IndexByte(path, ' ')
		if i < 0 {
			continue
		}

		dirPath, files := path[:i], path[i+1:]
		for _, file := range strings.Split(files, " ") {
			j := strings.IndexByte(file, '(')
			fileName, content := file[:j], file[j+1:len(file)-1]
			groups[content] = append(groups[content], dirPath+"/"+fileName)
		}
	}

	result := [][]string{}
	for _, group := range groups {
		if len(group) > 1 {
			result = append(result, group)
		}
	}

	return result
}
