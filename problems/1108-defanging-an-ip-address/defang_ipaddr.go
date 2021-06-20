package _108_defanging_an_ip_address

import "strings"

func defangIPaddr(address string) string {
	var sb strings.Builder
	sb.Grow(len(address) + 6)
	for _, r := range address {
		if r == '.' {
			sb.WriteString("[.]")
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
