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
		{"", "", []string{}},
		{"a:b", ":", []string{"a", "b"}},
		{"abc