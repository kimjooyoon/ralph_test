package domain

import (
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	result := EncodeBase64([]byte("hello"))
	if result != "aGVsbG8=" {
		t.Errorf("Expected \"aGVsbG8=\", got \"%s\"", result)
	}
}
