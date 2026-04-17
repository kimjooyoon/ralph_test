package domain

import "strings"

// Split splits s by sep.
//
// When sep is non-empty, behavior matches [strings.Split].
//
// When sep is empty: if s is empty, returns an empty slice. If s is non-empty, each UTF-8 byte
// of s becomes one element (not [strings.Split], which places each whole UTF-8 code sequence
// in a single element when sep is empty).
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	if sep == "" {
		result := make([]string, len(s))
		for i := 0; i < len(s); i++ {
			// Must slice by byte index — string(s[i]) would convert the byte to a Unicode
			// code point (e.g. 0xC3 → U+00C3 "Ã"), which breaks non-ASCII UTF-8.
			result[i] = s[i : i+1]
		}
		return result
	}
	return strings.Split(s, sep)
}
