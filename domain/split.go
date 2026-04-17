package domain

import "strings"

func Split(s, sep string) []string {
	var result []string
	current := ""
	for i := 0; i < len(s); {
		idx := strings.Index(s[i:], sep)
		if idx == -1 {
			current += s[i:]
			break
		}
		result = append(result, current)
		current = ""
		i += idx + len(sep)
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
