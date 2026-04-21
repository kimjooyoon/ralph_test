package domain

import (
	"testing"
)

func TestSplitSurrogatePairs(t *testing.T) {
	input := "\U0001F600" // Grinning Face emoji
	result := Split(input, "")
	if len(result) != 1 || result[0] != input {
		t.Errorf("Split(%q, \"\") = %v, want %v", input, result, []string{input})
	}
}
