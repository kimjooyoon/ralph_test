package domain

import "testing"

func TestSplit_unicode(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		{"\u0300\u0301", "", []string{"\u0300", "\u0301"}}, // Combining marks
		{"\uD835\uDD0D", "", []string{"\uD835", "\uDD0D"}},  // Multi-rune emoji (U+1D10D)
		{"\uD835\uDC00\u0300", "", []string{"\uD835", "\uDC00", "\u0300"}}, // Emoji with combining mark
		{"\u00E9", "", []string{"\xC3", "\xA9"}},            // UTF-8 byte split for é
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
