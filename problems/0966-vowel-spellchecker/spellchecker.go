package _966_vowel_spellchecker

import "strings"

func spellchecker(wordlist []string, queries []string) []string {
	exact := make(map[string]bool, len(wordlist))
	lower := make(map[string]string, len(wordlist))
	vowels := make(map[string]string, len(wordlist))
	replacer := strings.NewReplacer(
		"a", "_",
		"e", "_",
		"i", "_",
		"o", "_",
		"u", "_")
	for _, w := range wordlist {
		exact[w] = true

		l := strings.ToLower(w)
		if _, ok := lower[l]; !ok {
			lower[l] = w
		}

		v := replacer.Replace(l)
		if _, ok := vowels[v]; !ok {
			vowels[v] = w
		}
	}

	route := func(s string) string {
		if exact[s] {
			return s
		}
		l := strings.ToLower(s)
		if x, ok := lower[l]; ok {
			return x
		}
		v := replacer.Replace(l)
		if x, ok := vowels[v]; ok {
			return x
		}
		return ""
	}

	result := make([]string, 0, len(queries))
	for _, q := range queries {
		result = append(result, route(q))
	}

	return result
}
