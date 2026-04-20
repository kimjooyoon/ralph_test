package domain

import "testing"

func TestEncodeBase64(t *testing.T) {
    got := EncodeBase64("hello")
    want := "aGVsbG8="
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}
