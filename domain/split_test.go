package domain

import (
	"testing"
)

func TestSplitSurrogatePairs(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{
			input:    "\U0001D10D",
			sep:      "",
			expected: []string{"\U0001D10D"},
		},
		{
			input:    "abc\U0001D10Ddef",
			sep:      "",
			expected: []string{"a", "b", "c", "\U0001D10D", "d", "e", "f"},
		},
	}

	for _, test := range tests {
		result := Split(test.input, test.sep)
		if len(result) != len(test.expected) {
			t.Errorf("Split(%q, %q) returned %d elements, want %d", test.input, test.sep, len(result), len(test.expected))
			continue
		}
		for i := range result {
			if result[i] != test.expected[i] {
				t.Errorf("Split(%q, %q) element %d: got %q, want %q", test.input, test.sep, i, result[i], test.expected[i])
			}