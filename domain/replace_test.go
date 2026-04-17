package domain

import "testing"

func TestReplace(t *testing.T) {
	tests := []struct {
		input  string
		old    string
		new    string
		want   string
	}{
		{"hello", "l", "x", "hexxo"},
		{"aabbcc", "aa", "xx", "xxbbcc"},
		{"", "x", "y", ""},
		{"abc", "d", "x", "abc"},
		{"aaaa", "aa", "bb", "bbbb"},
	}
	for _, tt := range tests {
		if got := Replace(tt.input, tt.old, tt.new); got != tt.want {
			t.Errorf("Replace(%q, %q, %q) = %q, want %q", tt.input, tt.old, tt.new, got, tt.want)
		}
	}
}
