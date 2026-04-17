package domain

import "strings"

// Replace returns s with all occurrences of old replaced by new.
func Replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}
