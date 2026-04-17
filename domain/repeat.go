package domain

import "strings"

// Repeat returns s repeated n times.
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}
