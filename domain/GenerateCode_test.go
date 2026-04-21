package domain

import "testing"

func TestGenerateCode(t *testing.T) {
	t.Run("returns input as is", func(t *testing.T) {
		input := "for i in range(5): print(i)"
		result := GenerateCode(input)
		if result != input {
			t.Errorf("GenerateCode(%q) = %q, want %q", input, result, input)
		}
	})
}
