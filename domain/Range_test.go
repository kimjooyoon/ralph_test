package domain

import (
	"testing"
)

func TestRange(t *testing.T) {
	result := Range(1, 5)
	expected := []int{1, 2, 3, 4,