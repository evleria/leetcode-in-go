package _065_valid_number

import "regexp"

var isNumberRegexp = regexp.MustCompile(`^[+-]?((\d+\.?\d*)|\.\d+)([eE][+-]?\d+)?$`)

func isNumber(s string) bool {
	return isNumberRegexp.MatchString(s)
}
