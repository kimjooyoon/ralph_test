package domain

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abc", "cba"},
		{"", ""},
		{"a", "a"},
		{"hello world", "dlrow olleh"},
		{"12345", "54321"},
	}
	for _, tt := range tests {
		if got := Reverse(tt.input); got != tt.want {
			t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
