package domain

import "testing"

func TestSplit_spec_edges(t *testing.T) {
	if got := Split("", ""); len(got) != 0 {
		t.Errorf("Split(%q, %q) = %v, want empty slice", "", "", got)
	}
	if got := Split("a,b", ","); !equal(got, []string{"a", "b"}) {
		t.Errorf("Split(%q, %q) = %v, want [a b]", "a,b", ",", got)
	}
}

func TestSplit_unicode(t *testing.T) {
	tests := []struct {
		input string
		sep   string
		want  []string
	}{
		// UTF-8 byte split (spec): four bytes for two Latin-1 letters (À, Á).
		{"\xC3\x80\xC3\x81", "", []string{"\xC3", "\x80", "\xC3", "\x81"}},
		{"\xF0\x9D\x91\x8D", "", []string{"\xF0", "\x9D", "\x91", "\x8D"}}, // U+1D10D as four UTF-8 bytes
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
