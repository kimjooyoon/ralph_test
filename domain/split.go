package domain

import "strings"

// Split splits s by sep.
// When sep is empty and s is non-empty, each element is one UTF-8 byte of s (see specs/dsl.md).
// When sep is empty and s is empty, returns an empty slice (not [""]).
// Otherwise delegates to strings.Split.
func Split(s, sep string) []string {
	if sep == "" {
		if len(s) == 0 {
			return []string{}
		}
		out := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			out[i] = s[i : i+1]
		}
		return out
	}
	return strings.Split(s, sep)
}
