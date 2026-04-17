package domain

import "strings"

// Split splits s by sep (mostly strings.Split semantics, except Split("", sep) returns an empty slice).
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	if sep == "" {
		result := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			result[i] = string(s[i])
		}
		return result
	}
	return strings.Split(s, sep)
}
