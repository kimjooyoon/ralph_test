package domain

// Split splits s by sep and returns a slice of strings.
func Split(s, sep string) []string {
	var result []string
	current := ""
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			result = append(result, current)
			current = ""
		} else {
			current = current + string(s[i])
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}
