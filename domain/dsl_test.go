package domain

import "testing"

func TestEcho(t *testing.T) {
	if got := Echo("hi"); got != "hi" {
		t.Fatalf("Echo = %q, want %q", got, "hi")
	}
}

func TestHarnessPing(t *testing.T) {
	if got := Echo("ping"); got != "pong" {
		t.Fatalf("Echo(ping)=%q, want pong", got)
	}
}
