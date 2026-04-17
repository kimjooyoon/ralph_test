package domain

import "strings"

// Split splits s by sep (mostly strings.Split semantics, except Split("", sep) returns an empty slice).
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}
