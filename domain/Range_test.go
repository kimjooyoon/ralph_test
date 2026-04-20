package domain

import "testing"

func TestRangeNumericBounds(t *testing.T) {
    result := Range(1, 5)
    expected := []int{1, 2, 3, 4, 5}
    if len(result) != len(expected) {
        t.Errorf("Expected length %d, got %d", len(expected), len(result))
        return
    }
    for i := range result {