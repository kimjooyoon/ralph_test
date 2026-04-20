package domain

import (
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"", "x", []string{}},
		{"中文", "", []string{"中", "文"}},
		{"hello", " ", []string{"hello"}},
		{"abc123", "1", []string{"abc", "23"}),
	}

	for _, tt := range tests {
		result := Split(tt.input, tt.sep)
		if len(result) != len(tt.expected) {
			t.Errorf("Split(%q, %q) = %v, want %v", tt.input, tt.sep, result, tt.expected)
			continue
		}
		for i := range result {
			if result[i] != tt.expected[i] {
				t.Errorf("Split(%q, %q) = %v, want %v", tt.input, tt.sep, result, tt.expected)
				break
			}
		}
	}
}
