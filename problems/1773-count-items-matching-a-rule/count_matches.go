package _773_count_items_matching_a_rule

var mapping = map[string]int{
	"type":  0,
	"color": 1,
	"name":  2,
}

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	count := 0
	idx := mapping[ruleKey]
	for _, item := range items {
		if item[idx] == ruleValue {
			count++
		}
	}
	return count
}
