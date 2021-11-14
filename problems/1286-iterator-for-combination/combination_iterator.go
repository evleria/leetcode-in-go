package _286_iterator_for_combination

import (
	"sort"
)

type CombinationIterator struct {
	result []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	solution := solve(characters, combinationLength)
	return CombinationIterator{
		result: solution,
	}
}

func (c *CombinationIterator) Next() string {
	result := c.result[0]
	c.result = c.result[1:]
	return result
}

func (c *CombinationIterator) HasNext() bool {
	return len(c.result) > 0
}

func solve(characters string, length int) []string {
	solution := []string{}

	bytes := []byte(characters)
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	buffer := make([]byte, 0, length)
	backtrack(buffer, bytes, length, &solution)

	return solution
}

func backtrack(buffer []byte, characters []byte, length int, solution *[]string) {
	if len(buffer) == length {
		*solution = append(*solution, string(buffer))
		return
	}

	for i := 0; i < len(characters); i++ {
		backtrack(append(buffer, characters[i]), characters[i+1:], length, solution)
	}
}
