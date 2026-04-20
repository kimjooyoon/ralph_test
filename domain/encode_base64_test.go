package domain

import (
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "hello",
			expected: "aGVsbG8=",
		},
	}
	for _, tt := range tests {
		result := EncodeBase64([]byte(tt.input))
		if result != tt.expected {
			t.Errorf("EncodeBase64(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
