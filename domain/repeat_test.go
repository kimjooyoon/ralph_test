package domain

import "testing"

func TestRepeat(t *testing.T) {
	tests := []struct {
		input string
		n     int
		want  string
	}{
		{"a", 3, "aaa"},
		{"hello", 2, "hellohello"},
		{"", 5, ""},
		{"x", 0, ""},
		{"y", -1, ""},
	}
	for _, tt := range tests {
		if got := Repeat(tt.input, tt.n); got != tt.want {
			t.Errorf("Repeat(%q, %d) = %q, want %q", tt.input, tt.n, got, tt.want)
		}
	}
}
