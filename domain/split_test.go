package domain

import (
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input    string
		sep      string
		expected []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a:b:c", ":",