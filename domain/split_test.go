package domain

import "testing"

func TestSplit_unicode(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{"\xC3\x80\xC3\x81", "", []string{"\xC3\x80", "\xC3\x81"}}, // Combining marks
		{"\xF0\x9D\x91\x8D", "", []string{"\xF0", "\x9D", "\x91", "\x8D"}},  // Multi-rune emoji (U+1D10D)
		{"\xF0\x9D\x91\x8D\xE0\xA4\x82", "", []string{"\xF0", "\x9D", "\x91", "\x8D", "\xE0", "\xA4", "\x82"}}, // Emoji with combining mark
		{"\xC3\xA9", "", []string{"\xC3", "\xA9"}},            // UTF-8 byte split for é
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
