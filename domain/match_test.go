package domain

import (
	"testing"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		pattern string
		s       string
		want    bool
	}{
		{"^[a-z]+$", "abc123", false},
		{"^[0-9]+$", "12345", true},
		{"^hello$", "hello world", false},
		{"^hello$", "hello", true},
	}

	for _, tt := range tests {
		got := Match(tt.pattern, tt.s)
		if got != tt.want {
			t.Errorf("Match(%q, %q) = %v, want %v", tt.pattern, tt.s, got, tt.want)
		}
	}
}
