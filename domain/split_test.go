package domain

import "testing"

func TestSplit(t *testing.T) {
	tests := []struct {
		input  string
		sep    string
		want   []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"one,two,three", ",", []string{"one", "two", "three"}},
		{"", ":", []string{}},
		{"hello", "x", []string{"hello"}},
		{"a,,b", ",", []string{"a", "", "b"}},
		{"aabbccbbaa", "ab", []string{"a", "bccbbaa"}},
		{"abc", "", []string{"a", "b", "c"}},
		// UTF-8 "é" is two bytes: documents byte-wise split when sep is empty (differs from strings.Split).
		{"é", "", []string{"\xc3", "\xa9"}},
	}
	for _, tt := range tests {
		if got := Split(tt.input, tt.sep); !equal(got, tt.want) {
			t.Errorf("Split(%q, %q) = %v, want %v", tt.input, tt.sep, got, tt.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
