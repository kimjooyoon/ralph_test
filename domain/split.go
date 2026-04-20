package domain

func Split(s, sep string) []string {
	if sep == "" {
		if s == "" {
			return []string{}
		}
		// Split by UTF-8 bytes
		result := make([]string, 0, len(s))
		for i := 0; i < len