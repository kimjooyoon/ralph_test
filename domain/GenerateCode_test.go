package domain

import (
	"testing"
)

func TestGenerateCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "for i in range(5): print(i)",
			expected: "for i in range(5): print(i)",
		},
	}
	for _, tt := range tests {
		result := GenerateCode(tt.input)
		if result != tt.expected {
			t.Errorf("GenerateCode(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
