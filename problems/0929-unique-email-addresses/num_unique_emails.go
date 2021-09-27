package _929_unique_email_addresses

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for _, email := range emails {
		atIdx := strings.IndexByte(email, '@')
		local, domain := email[:atIdx], email[atIdx+1:]
		if plusIdx := strings.IndexByte(local, '+'); plusIdx != -1 {
			local = local[:plusIdx]
		}
		local = strings.ReplaceAll(local, ".", "")
		email = fmt.Sprintf("%s@%s", local, domain)
		m[email] = true
	}

	return len(m)
}
