package domain

import "testing"

func TestJoin(t *testing.T) {
	if got := Join("a", "b"); got != "a:b" {
		t.Fatalf("Join = %q, want a:b", got)
	}
}
