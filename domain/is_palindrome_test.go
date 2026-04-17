package domain

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcba", true},
		{"abc", false},
		{"", true},
		{"a", true},
		{"12321", true},
		{"1234", false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.input); got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
