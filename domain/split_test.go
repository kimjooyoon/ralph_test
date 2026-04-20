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
		{"a", "a", []string{""}},
		{"a", "b", []string{"a"}},
		{"中文", "", []string{"中", "文"}},
	}

	for _, tt := range tests {
		result := Split(tt.input, tt.sep)
		if len(result) != len(tt.expected) {
			t.Errorf("Split(%q, %q) = %v, want %v", tt.input, tt.sep, result, tt.expected)
			continue
		}
		for i := range result {