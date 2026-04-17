package domain

import "testing"

func TestTrim(t *testing.T) {
	if got := Trim("  abc  "); got != "abc" {
		t.Errorf("Trim = %q, want %q", got, "abc")
	}
}
